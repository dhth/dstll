package server

import (
	"bytes"
	"context"
	"embed"
	"errors"
	"fmt"
	htemplate "html/template"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/dhth/dstll/internal/utils"
	"github.com/dhth/dstll/tsutils"
)

const (
	startPort               = 8100
	endPort                 = 8500
	minResultsForLoadingBar = 400
	resultsPerPage          = 10
)

const (
	backgroundColor = "#1f1f24"
	filepathColor   = "#ffd166"
	headerColor     = "#f15bb5"
	navigationColor = "#00bbf9"
	activePageColor = "#f15bb5"
)

var ErrNoPortOpen = errors.New("No open port found")

//go:embed html
var tplFolder embed.FS

type htmlData struct {
	Results         map[string][]htemplate.HTML
	Pages           []int
	CurrentPage     int
	NumFiles        int
	BackgroundColor string
	HeaderColor     string
	FilepathColor   string
	NavigationColor string
	ActivePageColor string
}

func serveResults(files []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		page := r.PathValue("page")
		pageNum := 1
		var err error

		if page != "" {
			pageNum, err = strconv.Atoi(page)
			if err != nil {
				log.Printf("Got bad page: %v", page)
				http.Error(w, "Bad page", http.StatusBadRequest)
				return
			}
		}

		startIndex, endIndex, err := utils.GetIndexRange(pageNum, len(files), resultsPerPage)
		if err != nil {
			http.Error(w, fmt.Sprintf("Bad page: %q", err), http.StatusBadRequest)
			return
		}

		results := getResults(files[startIndex:endIndex])
		numPages := len(files) / resultsPerPage

		if len(files)%resultsPerPage != 0 {
			numPages++
		}

		pages := make([]int, numPages)

		for i := range numPages {
			pages[i] = i + 1
		}

		res1 := htmlData{
			Results:         results,
			Pages:           pages,
			CurrentPage:     pageNum,
			NumFiles:        len(results),
			BackgroundColor: backgroundColor,
			HeaderColor:     headerColor,
			FilepathColor:   filepathColor,
			NavigationColor: navigationColor,
			ActivePageColor: activePageColor,
		}

		w.Header().Add("Server", "Go")

		files := []string{
			"html/base.tmpl",
			"html/partials/nav.tmpl",
			"html/pages/home.tmpl",
		}

		ts, err := template.ParseFS(tplFolder, files...)
		if err != nil {
			log.Printf("Error getting template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", res1)
		if err != nil {
			log.Printf("Error executing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func getResults(fPaths []string) map[string][]htemplate.HTML {
	resultsChan := make(chan tsutils.Result)
	results := make(map[string][]htemplate.HTML)

	for _, fPath := range fPaths {
		go tsutils.GetLayout(resultsChan, fPath)
	}

	for range fPaths {
		r := <-resultsChan
		if r.Err != nil {
			continue
		}

		if len(r.Results) == 0 {
			continue
		}

		htmlResults := make([]htemplate.HTML, len(r.Results))

		elementsCombined := strings.Join(r.Results, "\n\n")
		var b bytes.Buffer

		err := quick.Highlight(&b, elementsCombined, r.FPath, "html", "xcode-dark")
		if err != nil {
			htmlResults = append(htmlResults, htemplate.HTML(fmt.Sprintf("<p>%s</p>", elementsCombined)))
		} else {
			htmlResults = append(htmlResults, htemplate.HTML(b.String()))
		}
		results[r.FPath] = htmlResults
	}

	return results
}

func Start(fPaths []string) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", serveResults(fPaths))
	mux.HandleFunc("GET /page/{page}", serveResults(fPaths))

	port, err := findOpenPort(startPort, endPort)
	if err != nil {
		fmt.Printf("Couldn't find an open port between %d-%d", startPort, endPort)
	}
	server := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%d", port),
		Handler: mux,
	}

	go func() {
		fmt.Printf("Starting server. Open http://%s/ in your browser.\n", server.Addr)
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Error running server: %q", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutDownCtx, shutDownRelease := context.WithTimeout(context.Background(), time.Second*3)
	defer shutDownRelease()

	err = server.Shutdown(shutDownCtx)
	if err != nil {
		fmt.Printf("Error shutting down: %v\nTrying forceful shutdown\n", err)
		closeErr := server.Close()
		if closeErr != nil {
			fmt.Printf("Forceful shutdown failed: %v\n", closeErr)
		} else {
			fmt.Printf("Forceful shutdown successful\n")
		}
	}
	fmt.Printf("\nbye 👋\n")
}

func findOpenPort(startPort, endPort int) (int, error) {
	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("127.0.0.1:%d", port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			defer listener.Close()
			return port, nil
		}
	}
	return 0, ErrNoPortOpen
}

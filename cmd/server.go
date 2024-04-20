package cmd

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
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/alecthomas/chroma/quick"
	"github.com/dhth/dstll/tsutils"
)

const (
	startPort = 8100
	endPort   = 8500
)

const (
	backgroundColor = "#1f1f24"
	HeaderColor     = "#41a1c0"
	filepathColor   = "#ffd166"
	headerColor     = "#41a1c0"
)

var (
	noPortOpenErr = errors.New("No open port found")
)

//go:embed web
var tplFolder embed.FS

type htmlResults struct {
	Results         map[string][]htemplate.HTML
	BackgroundColor string
	HeaderColor     string
	FilepathColor   string
}

func (res htmlResults) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"web/html/base.tmpl",
		"web/html/pages/home.tmpl",
	}

	ts, err := template.ParseFS(tplFolder, files...)
	if err != nil {
		log.Printf("Error getting template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", res)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func spinner(delay time.Duration, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\rfetching results %c", r)
				time.Sleep(delay)
			}
		}
	}
}

func getResults(fPaths []string) map[string][]htemplate.HTML {
	done := make(chan struct{})
	go spinner(time.Millisecond*100, done)

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
	done <- struct{}{}
	return results
}

func startServer(fPaths []string) {
	results := getResults(fPaths)

	res := htmlResults{
		Results:         results,
		BackgroundColor: backgroundColor,
		HeaderColor:     headerColor,
		FilepathColor:   filepathColor,
	}

	if len(res.Results) == 0 {
		return
	}

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", res)

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
	return 0, noPortOpenErr
}

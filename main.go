package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	trimPrefix = flag.String("trim-prefix", "", "prefix to trim from the file path")
)

var (
	FilePathIncorrectErr = errors.New("file path incorrect")
	FileNameIncorrectErr = errors.New("file name incorrect")
)

func getLayout(resultsChan chan<- Result, filePath string) {
	ext, err := getFileExtension(filePath)
	if err != nil {
		resultsChan <- Result{fPath: filePath, err: err}
		return
	}

	fContent, err := os.ReadFile(filePath)
	if err != nil {
		resultsChan <- Result{fPath: filePath, err: err}
		return
	}

	var elements []string
	switch ext {
	case FTScala:
		elements, err = getScalaData(fContent)
	case FTGo:
		elements, err = getGoData(fContent)
	case FTPython:
		elements, err = getPyData(fContent)
	default:
		return
	}
	if err != nil {
		resultsChan <- Result{fPath: filePath, err: err}
	} else {
		resultsChan <- Result{fPath: filePath, results: elements}
	}
}

func main() {
	flag.Parse()

	var fPaths []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fPaths = append(fPaths, scanner.Text())
	}

	resultsChan := make(chan Result)
	results := make(map[string][]string)

	for _, fPath := range fPaths {
		go getLayout(resultsChan, fPath)
	}

	for range fPaths {
		r := <-resultsChan
		if r.err == nil {
			results[r.fPath] = r.results
		}
	}

	for _, fPath := range fPaths {
		v, ok := results[fPath]
		if !ok {
			continue
		}

		if len(v) == 0 {
			continue
		}
		if *trimPrefix != "" {
			fmt.Println("👉 " + filePathStyle.Render(strings.TrimPrefix(fPath, *trimPrefix)))
		} else {
			fmt.Println("👉 " + filePathStyle.Render(fPath))
		}
		fmt.Println()

		for _, elem := range v {
			fmt.Println(tsElementStyle.Render(elem))
		}
		fmt.Print("\n")
		fmt.Print(dividerStyle.Render(strings.Repeat(".", 80)))
		fmt.Print("\n\n")
	}
}

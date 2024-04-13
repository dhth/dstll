package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dhth/dstll/tsutils"
)

func ShowResults(trimPrefix string, plain bool) {

	var fPaths []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fPaths = append(fPaths, scanner.Text())
	}

	resultsChan := make(chan tsutils.Result)
	results := make(map[string][]string)

	for _, fPath := range fPaths {
		go tsutils.GetLayout(resultsChan, fPath)
	}

	for range fPaths {
		r := <-resultsChan
		if r.Err == nil {
			results[r.FPath] = r.Results
		}
	}

	switch plain {
	case true:
		printPlainOutput(fPaths, results, trimPrefix)
	case false:
		printColorOutput(fPaths, results, trimPrefix)
	}
}

func printColorOutput(fPaths []string, results map[string][]string, trimPrefix string) {
	for _, fPath := range fPaths {
		v, ok := results[fPath]
		if !ok {
			continue
		}

		if len(v) == 0 {
			continue
		}
		if trimPrefix != "" {
			fmt.Println("👉 " + filePathStyle.Render(strings.TrimPrefix(fPath, trimPrefix)))
		} else {
			fmt.Println("👉 " + filePathStyle.Render(fPath))
		}
		fmt.Println()

		for _, elem := range v {
			fmt.Println(tsElementStyle.Render(elem))
			fmt.Println()
		}
		fmt.Print(dividerStyle.Render(strings.Repeat(".", 80)))
		fmt.Print("\n\n")
	}
}

func printPlainOutput(fPaths []string, results map[string][]string, trimPrefix string) {
	for _, fPath := range fPaths {
		v, ok := results[fPath]
		if !ok {
			continue
		}

		if len(v) == 0 {
			continue
		}
		if trimPrefix != "" {
			fmt.Println("-> " + strings.TrimPrefix(fPath, trimPrefix))
		} else {
			fmt.Println("-> " + fPath)
		}
		fmt.Println()

		for _, elem := range v {
			fmt.Println(elem)
			fmt.Println()
		}
		fmt.Print(strings.Repeat(".", 80))
		fmt.Print("\n\n")
	}
}

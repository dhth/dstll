package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhth/dstll/tsutils"
)

type writeResult struct {
	path string
	err  error
}

func ShowResults(results []tsutils.Result, trimPrefix string, plain bool) {
	switch plain {
	case true:
		printPlainOutput(results, trimPrefix)
	case false:
		printColorOutput(results, trimPrefix)
	}
}

func printColorOutput(results []tsutils.Result, trimPrefix string) {
	for i, result := range results {
		if result.Err != nil {
			continue
		}

		if trimPrefix != "" {
			fmt.Println("👉 " + filePathStyle.Render(strings.TrimPrefix(result.FPath, trimPrefix)))
		} else {
			fmt.Println("👉 " + filePathStyle.Render(result.FPath))
		}
		fmt.Println()

		var r []string
		for _, elem := range result.Results {
			r = append(r, tsElementStyle.Render(elem))
		}
		fmt.Println(strings.Join(r, "\n\n"))

		if i < len(results)-1 {
			fmt.Printf("\n%s\n\n", dividerStyle.Render(strings.Repeat(".", 80)))
		}
	}
}

func printPlainOutput(results []tsutils.Result, trimPrefix string) {
	for i, result := range results {
		if result.Err != nil {
			continue
		}

		if trimPrefix != "" {
			fmt.Println("-> " + strings.TrimPrefix(result.FPath, trimPrefix))
		} else {
			fmt.Println("-> " + result.FPath)
		}
		fmt.Println()

		fmt.Println(strings.Join(result.Results, "\n\n"))

		if i < len(results)-1 {
			fmt.Printf("\n%s\n\n", strings.Repeat(".", 80))
		}
	}
}

func writeToFile(resultsChan chan<- writeResult, path string, contents []string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		resultsChan <- writeResult{path, err}
		return
	}

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		resultsChan <- writeResult{path, err}
		return
	}
	defer file.Close()

	_, err = file.WriteString(strings.Join(contents, "\n") + "\n")
	resultsChan <- writeResult{path, err}
}

func WriteResults(results []tsutils.Result, outDir string, quiet bool) {
	resultsChan := make(chan writeResult)
	var successes []string
	errors := make(map[string]error)

	counter := 0
	for _, result := range results {
		outPath := filepath.Join(outDir, result.FPath)
		if len(result.Results) > 0 {
			go writeToFile(resultsChan, outPath, result.Results)
			counter++
		}
	}

	for i := 0; i < counter; i++ {
		r := <-resultsChan
		if r.err != nil {
			errors[r.path] = r.err
		} else {
			successes = append(successes, r.path)
		}
	}

	errorList := make([]string, len(errors))
	c := 0
	for p, e := range errors {
		errorList[c] = fmt.Sprintf("%s: %s", p, e.Error())
		c++
	}

	if !quiet {
		if len(successes) > 0 {
			fmt.Printf("The following files were written:\n%s\n", strings.Join(successes, "\n"))
		}
	}

	if len(errorList) > 0 {
		if !quiet {
			if len(successes) > 0 {
				fmt.Print("\n---\n\n")
			}
		}
		fmt.Fprintf(os.Stderr, "The following errors were encountered:\n%s\n", strings.Join(errorList, "\n"))
	}
}

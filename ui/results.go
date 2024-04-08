package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dhth/layitout/tsutils"
)

func ShowResults(trimPrefix string) {

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

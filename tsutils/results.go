package tsutils

func GetResults(fPaths []string) []Result {
	resultsChan := make(chan Result)
	results := make([]Result, len(fPaths))

	for _, fPath := range fPaths {
		go GetLayout(resultsChan, fPath)
	}

	for i := range fPaths {
		r := <-resultsChan
		results[i] = r
	}

	return results
}

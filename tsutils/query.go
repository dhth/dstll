package tsutils

import "os"

func GetLayout(resultsChan chan<- Result, filePath string) {
	ext, err := getFileExtension(filePath)
	if err != nil {
		resultsChan <- Result{FPath: filePath, Err: err}
		return
	}

	fContent, err := os.ReadFile(filePath)
	if err != nil {
		resultsChan <- Result{FPath: filePath, Err: err}
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
		resultsChan <- Result{FPath: filePath, Err: err}
	} else {
		resultsChan <- Result{FPath: filePath, Results: elements}
	}
}

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
		objectChan := make(chan Result)
		classChan := make(chan Result)
		fnChan := make(chan Result)

		chans := []chan Result{objectChan, classChan, fnChan}
		go getScalaObjects(objectChan, fContent)
		go getScalaClasses(classChan, fContent)
		go getScalaFunctions(fnChan, fContent)

		for _, ch := range chans {
			r := <-ch
			if r.Err == nil {
				for _, elem := range r.Results {
					elements = append(elements, elem)
				}
			}
		}
	case FTGo:
		fnChan := make(chan Result)
		methodChan := make(chan Result)
		chans := []chan Result{fnChan, methodChan}

		go getGoFuncs(fnChan, fContent)
		go getGoMethods(methodChan, fContent)

		for _, ch := range chans {
			r := <-ch
			if r.Err == nil {
				for _, elem := range r.Results {
					elements = append(elements, elem)
				}
			}
		}
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

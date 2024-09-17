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
	case FTGo:
		typeChan := make(chan Result)
		fnChan := make(chan Result)
		methodChan := make(chan Result)
		chans := []chan Result{typeChan, fnChan, methodChan}

		go getGoTypes(typeChan, fContent)
		go getGoFuncs(fnChan, fContent)
		go getGoMethods(methodChan, fContent)

		for _, ch := range chans {
			r := <-ch
			if r.Err == nil {
				elements = append(elements, r.Results...)
			}
		}
	case FTPython:
		elements, err = getPyData(fContent)
	case FTRust:
		structChan := make(chan Result)
		enumChan := make(chan Result)
		traitChan := make(chan Result)
		fnChan := make(chan Result)
		chans := []chan Result{structChan, enumChan, traitChan, fnChan}

		go getRustStructs(structChan, fContent)
		go getRustEnums(enumChan, fContent)
		go getRustTraits(traitChan, fContent)
		go getRustFuncs(fnChan, fContent)

		for _, ch := range chans {
			r := <-ch
			if r.Err == nil {
				elements = append(elements, r.Results...)
			}
		}
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
				elements = append(elements, r.Results...)
			}
		}
	default:
		return
	}
	if err != nil {
		resultsChan <- Result{FPath: filePath, Err: err}
	} else {
		resultsChan <- Result{FPath: filePath, Results: elements}
	}
}

package main

import "strings"

func getFileExtension(filePath string) (FileType, error) {
	fPathEls := strings.Split(filePath, ".")
	if len(fPathEls) < 2 {
		return FTNone, FilePathIncorrectErr
	}
	var ft FileType
	switch fPathEls[len(fPathEls)-1] {
	case "scala":
		ft = FTScala
	case "go":
		ft = FTGo
	case "py":
		ft = FTPython
	default:
		return FTNone, FilePathIncorrectErr
	}

	fNameEls := strings.Split(filePath, "/")
	if strings.Split(fNameEls[len(fNameEls)-1], ".")[0] == "" {
		return FTNone, FileNameIncorrectErr
	}

	return ft, nil
}

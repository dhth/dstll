package tsutils

import (
	"errors"
	"strings"
)

var (
	ErrFilePathIncorrect = errors.New("file path incorrect")
	ErrFileNameIncorrect = errors.New("file name incorrect")
)

func getFileExtension(filePath string) (FileType, error) {
	fPathEls := strings.Split(filePath, ".")
	if len(fPathEls) < 2 {
		return FTNone, ErrFilePathIncorrect
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
		return FTNone, ErrFilePathIncorrect
	}

	fNameEls := strings.Split(filePath, "/")
	if strings.Split(fNameEls[len(fNameEls)-1], ".")[0] == "" {
		return FTNone, ErrFileNameIncorrect
	}

	return ft, nil
}

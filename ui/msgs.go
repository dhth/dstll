package ui

import "github.com/dhth/dstll/tsutils"

type hideHelpMsg struct{}

type FileRead struct {
	contents string
	err      error
}

type FileResultsReceivedMsg struct {
	result tsutils.Result
}

type ViewFileFinishedmsg struct {
	filePath string
	err      error
}

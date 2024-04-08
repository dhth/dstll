package ui

import "github.com/dhth/layitout/tsutils"

type HideHelpMsg struct{}

type FileRead struct {
	contents string
	err      error
}

type FileResultsReceivedMsg struct {
	result tsutils.Result
}

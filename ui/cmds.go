package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dhth/dstll/tsutils"
)

func chooseFile(filePath string) tea.Cmd {
	return func() tea.Msg {
		fContent, err := os.ReadFile(filePath)
		if err != nil {
			return FileRead{err: err}
		}
		return FileRead{contents: string(fContent)}
	}
}

func getFileResults(filePath string) tea.Cmd {
	return func() tea.Msg {
		resultsChan := make(chan tsutils.Result)
		go tsutils.GetLayout(resultsChan, filePath)

		result := <-resultsChan

		return FileResultsReceivedMsg{result}
	}
}

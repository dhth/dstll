package ui

import (
	"os"
	"os/exec"
	"time"

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

func openFile(filePath string, cmd []string) tea.Cmd {
	openCmd := append(cmd, filePath)
	c := exec.Command(openCmd[0], openCmd[1:]...)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return ViewFileFinishedmsg{filePath, err}
	})
}

func getFileResults(filePath string) tea.Cmd {
	return func() tea.Msg {
		resultsChan := make(chan tsutils.Result)
		go tsutils.GetLayout(resultsChan, filePath)

		result := <-resultsChan

		return FileResultsReceivedMsg{result}
	}
}

func hideHelp(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(time.Time) tea.Msg {
		return hideHelpMsg{}
	})
}

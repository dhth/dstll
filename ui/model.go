package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dhth/dstll/filepicker"
)

type Pane uint

const (
	fileExplorerPane Pane = iota
	resultPane
)

type model struct {
	config                Config
	resultVP              viewport.Model
	resultVPReady         bool
	resultsCache          map[string]string
	filepicker            filepicker.Model
	selectedFile          string
	quitting              bool
	activePane            Pane
	terminalHeight        int
	terminalWidth         int
	message               string
	showHelp              bool
	noConstructsMsg       string
	supportedFileTypes    []string
	unsupportedFileMsg    string
	fileExplorerPaneWidth int
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		hideHelp(time.Minute*1),
		m.filepicker.Init(),
	)
}

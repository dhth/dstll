package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dhth/layitout/filepicker"
	"github.com/dhth/layitout/tsutils"
)

type Pane uint

const (
	fileExplorerPane Pane = iota
	resultPane
)

type model struct {
	resultVP              viewport.Model
	resultVPReady         bool
	resultsCache          map[string]tsutils.Result
	filepicker            filepicker.Model
	selectedFile          string
	quitting              bool
	err                   error
	activePane            Pane
	lastPane              Pane
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
	return m.filepicker.Init()
}

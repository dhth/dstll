package ui

import (
	"bytes"

	"github.com/alecthomas/chroma/quick"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const useHighPerformanceRenderer = false

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	m.message = ""

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "tab", "shift+tab":
			if m.activePane == fileExplorerPane {
				m.activePane = resultPane
			} else {
				m.activePane = fileExplorerPane
				m.resultVP.GotoTop()
			}
		case "o":
			if m.config.ViewFileCmd != nil {
				if m.activePane == fileExplorerPane {
					if m.filepicker.IsCurrentAFile {
						cmds = append(cmds, openFile(m.filepicker.Current, *m.config.ViewFileCmd))
					}
				}
			} else {
				m.message = "you haven't configured view_file_command, run dstll -help to learn more"
			}
		}
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.filepicker.Height = msg.Height - 8

		if !m.resultVPReady {
			m.resultVP = viewport.New(msg.Width-m.fileExplorerPaneWidth-10, msg.Height-8)
			m.resultVP.HighPerformanceRendering = useHighPerformanceRenderer
			m.resultVPReady = true
		} else {
			m.resultVP.Width = msg.Width - m.fileExplorerPaneWidth - 10
			m.resultVP.Height = msg.Height - 8
		}
	case hideHelpMsg:
		m.showHelp = false
	case FileRead:
		if msg.err != nil {
			m.message = msg.err.Error()
		} else {
			m.resultVP.SetContent(msg.contents)
		}
	case FileResultsReceivedMsg:
		if msg.result.Err != nil {
			m.message = msg.result.Err.Error()
		} else {
			if len(msg.result.Results) == 0 {
				m.resultVP.SetContent(m.noConstructsMsg)
				m.resultsCache[msg.result.FPath] = m.noConstructsMsg
			} else {
				s := "👉 " + filePathStyleTUI.Render(msg.result.FPath) + "\n\n"
				for _, elem := range msg.result.Results {
					var b bytes.Buffer
					err := quick.Highlight(&b, elem, msg.result.FPath, "terminal16m", "xcode-dark")
					if err != nil {
						s += tsElementStyle.Render(elem)
					} else {
						s += b.String()
					}
					s += "\n\n"
				}
				m.resultVP.SetContent(s)
				m.resultsCache[msg.result.FPath] = s
			}
		}
	}

	var cmd tea.Cmd
	switch m.activePane {
	case fileExplorerPane:
		m.filepicker, cmd = m.filepicker.Update(msg)
		cmds = append(cmds, cmd)
		if m.filepicker.CanSelect(m.filepicker.Current) {
			m.selectedFile = m.filepicker.Current
			resultFromCache, ok := m.resultsCache[m.filepicker.Current]
			if !ok {
				cmds = append(cmds, getFileResults(m.filepicker.Current))
			} else {
				m.resultVP.SetContent(resultFromCache)
			}
		} else {
			m.resultVP.SetContent(unsupportedFileStyle.Render(m.unsupportedFileMsg))
		}
	case resultPane:
		m.resultVP, cmd = m.resultVP.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

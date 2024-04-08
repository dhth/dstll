package ui

import (
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
			m.resultsCache[msg.result.FPath] = msg.result

			if len(msg.result.Results) == 0 {
				m.resultVP.SetContent(m.noConstructsMsg)
			} else {
				s := "👉 " + msg.result.FPath + "\n\n"
				for _, elem := range msg.result.Results {
					s += tsElementStyle.Render(elem)
					s += "\n\n"
				}
				m.resultVP.SetContent(s)
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
				if resultFromCache.Err != nil {
					m.message = resultFromCache.Err.Error()
				} else {
					s := "👉 " + filePathStyle.Render(m.filepicker.Current) + "\n\n"
					if len(resultFromCache.Results) == 0 {
						m.resultVP.SetContent(s + noConstructsStyle.Render(m.noConstructsMsg))
					} else {
						for _, elem := range resultFromCache.Results {
							s += tsElementStyle.Render(elem)
							s += "\n\n"
						}
						m.resultVP.SetContent(s)
					}
				}
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

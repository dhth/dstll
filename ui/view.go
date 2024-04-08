package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var content string
	var footer string

	var statusBar string
	if m.message != "" {
		statusBar = RightPadTrim(m.message, m.terminalWidth)
	}

	fileExplorerStyle.GetWidth()
	switch m.activePane {
	case fileExplorerPane:
		fExplorer := lipgloss.JoinVertical(lipgloss.Left, "\n  "+activePaneHeaderStyle.Render("Files")+"\n\n"+fileExplorerStyle.Render(m.filepicker.View()))
		resultView := lipgloss.JoinVertical(lipgloss.Left, "\n"+inActivePaneHeaderStyle.Render("Results")+"\n\n"+m.resultVP.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top, fExplorer, resultView)
	case resultPane:
		fExplorer := lipgloss.JoinVertical(lipgloss.Left, "\n  "+inActivePaneHeaderStyle.Render("Files")+"\n\n"+fileExplorerStyle.Render(m.filepicker.View()))
		resultView := lipgloss.JoinVertical(lipgloss.Left, "\n"+activePaneHeaderStyle.Render("Results")+"\n\n"+m.resultVP.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top, fExplorer, resultView)
	}

	var cwdBar string
	if m.filepicker.CurrentDirectory != "." {
		cwdBar = cwdStyle.Render(fmt.Sprintf("cwd: %s", m.filepicker.CurrentDirectory))
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(DefaultForegroundColor)).
		Background(lipgloss.Color("#7c6f64"))

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("Press ? for help")
	}

	footerStr := fmt.Sprintf("%s%s",
		modeStyle.Render("layitout"),
		helpMsg,
	)
	footer = footerStyle.Render(footerStr)

	return lipgloss.JoinVertical(lipgloss.Left,
		content,
		cwdBar,
		statusBar,
		footer,
	)
}

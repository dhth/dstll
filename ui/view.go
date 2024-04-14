package ui

import (
	"fmt"
	"strings"

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

	// It seems that using terminal color codes in the viewport is leaking some
	// information into the file picker pane, resulting in some lines in it
	// being colored under some circumstances. This color reset pane fixes that.
	// Not the most elegant solution, but seems to be doing the job.
	var colorResetPane string
	if m.resultVPReady {
		colorResetPane = strings.Repeat("\033[0m\n", m.filepicker.Height-1)
		colorResetPane += "\033[0m"
	}

	fileExplorerStyle.GetWidth()
	switch m.activePane {
	case fileExplorerPane:
		fExplorer := lipgloss.JoinVertical(lipgloss.Left, "\n  "+activePaneHeaderStyle.Render("Files")+"\n\n"+fileExplorerStyle.Render(m.filepicker.View()))
		resultView := lipgloss.JoinVertical(lipgloss.Left, "\n"+inActivePaneHeaderStyle.Render("Results")+"\n\n"+m.resultVP.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top, colorResetPane, fExplorer, resultView)
	case resultPane:
		fExplorer := lipgloss.JoinVertical(lipgloss.Left, "\n  "+inActivePaneHeaderStyle.Render("Files")+"\n\n"+fileExplorerStyle.Render(m.filepicker.View()))
		resultView := lipgloss.JoinVertical(lipgloss.Left, "\n"+activePaneHeaderStyle.Render("Results")+"\n\n"+m.resultVP.View())
		content = lipgloss.JoinHorizontal(lipgloss.Top, colorResetPane, fExplorer, resultView)
	}

	var cwdBar string
	if m.filepicker.CurrentDirectory != "." {
		cwdBar = cwdStyle.Render(fmt.Sprintf("cwd: %s", m.filepicker.CurrentDirectory))
	}

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(DefaultForegroundColor)).
		Background(lipgloss.Color(FooterColor))

	var helpMsg string
	if m.showHelp {
		helpMsg = " " + helpMsgStyle.Render("Press ? for help")
	}

	footerStr := fmt.Sprintf("%s%s",
		modeStyle.Render("dstll"),
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

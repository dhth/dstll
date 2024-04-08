package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	ActiveHeaderColor      = "#fe8019"
	ResultsHeaderColor     = "#b8bb26"
	InactivePaneColor      = "#928374"
	DisabledFileColor      = "#928374"
	DirectoryColor         = "#83a598"
	NoConstructsColor      = "#fabd2f"
	UnsupportedFileColor   = "#928374"
	CWDColor               = "#928374"
	FilepathColor          = "#8ec07c"
	TSElementColor         = "#d3869b"
	DividerColor           = "#665c54"
	DefaultForegroundColor = "#282828"
)

var (
	filePathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(FilepathColor))

	tsElementStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(TSElementColor))

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(DividerColor))
	baseStyle = lipgloss.NewStyle().
			PaddingLeft(1).
			PaddingRight(1).
			Foreground(lipgloss.Color(DefaultForegroundColor))

	fileExplorerStyle = lipgloss.NewStyle().
				Width(45).
				PaddingRight(2).
				PaddingBottom(1)

	activePaneHeaderStyle = baseStyle.Copy().
				Align(lipgloss.Left).
				Bold(true).
				Background(lipgloss.Color(ActiveHeaderColor))

	inActivePaneHeaderStyle = activePaneHeaderStyle.Copy().
				Background(lipgloss.Color(InactivePaneColor))

	unsupportedFileStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(UnsupportedFileColor))

	noConstructsStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(NoConstructsColor))

	cwdStyle = baseStyle.Copy().
			PaddingRight(0).
			Foreground(lipgloss.Color(CWDColor))

	modeStyle = baseStyle.Copy().
			Align(lipgloss.Center).
			Bold(true).
			Background(lipgloss.Color("#b8bb26"))

	helpMsgStyle = baseStyle.Copy().
			Bold(true).
			Foreground(lipgloss.Color("#83a598"))
)

package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	ActiveHeaderColor      = "#fd7474"
	InactivePaneColor      = "#c1abea"
	DisabledFileColor      = "#8a93a5"
	DirectoryColor         = "#76a9f9"
	NoConstructsColor      = "#fabd2f"
	UnsupportedFileColor   = "#928374"
	CWDColor               = "#c1abea"
	FilepathColor          = "#fd7474"
	TSElementColor         = "#76a9f9"
	DividerColor           = "#665c54"
	DefaultForegroundColor = "#282828"
	ModeColor              = "#76a9f9"
	HelpMsgColor           = "#83a598"
	FooterColor            = "#7c6f64"
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
			Background(lipgloss.Color(ModeColor))

	helpMsgStyle = baseStyle.Copy().
			Bold(true).
			Foreground(lipgloss.Color(HelpMsgColor))
)

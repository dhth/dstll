package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	ActiveHeaderColor      = "#fc5fa3"
	InactivePaneColor      = "#d0a8ff"
	DisabledFileColor      = "#6c7986"
	DirectoryColor         = "#41a1c0"
	NoConstructsColor      = "#fabd2f"
	UnsupportedFileColor   = "#928374"
	CWDColor               = "#d0a8ff"
	FilepathColor          = "#fc5fa3"
	FilepathColorTUI       = "#d0a8ff"
	TSElementColor         = "#41a1c0"
	DividerColor           = "#6c7986"
	DefaultForegroundColor = "#282828"
	ModeColor              = "#fc5fa3"
	HelpMsgColor           = "#83a598"
	FooterColor            = "#7c6f64"
)

var (
	filePathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(FilepathColor))

	filePathStyleTUI = lipgloss.NewStyle().
				Foreground(lipgloss.Color(FilepathColorTUI))

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

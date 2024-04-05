package main

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	FilepathColor  = "#8ec07c"
	TSElementColor = "#d3869b"
	DividerColor   = "#665c54"
)

var (
	filePathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(FilepathColor))

	tsElementStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(TSElementColor))

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(DividerColor))
)

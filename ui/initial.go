package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/dstll/filepicker"
)

func InitialModel(config Config) Model {
	fp := filepicker.New()
	supportedFT := []string{
		".go",
		".py",
		".rs",
		".scala",
	}

	var unsupportedFTMsg strings.Builder
	unsupportedFTMsg.WriteString("dstll will show constructs for the following file types:\n")
	for _, ft := range supportedFT {
		fmt.Fprintf(&unsupportedFTMsg, "%s\n", ft)
	}

	fpWidth := 40

	fp.AllowedTypes = supportedFT
	fp.AutoHeight = false
	fp.Width = fpWidth
	fp.Styles.Selected = fp.Styles.Selected.Foreground(lipgloss.Color(ActiveHeaderColor))
	fp.Styles.Cursor = fp.Styles.Cursor.Foreground(lipgloss.Color(ActiveHeaderColor))
	fp.Styles.DisabledFile = fp.Styles.DisabledFile.Foreground(lipgloss.Color(DisabledFileColor))
	fp.Styles.Directory = fp.Styles.Directory.Foreground(lipgloss.Color(DirectoryColor))

	m := Model{
		config:                config,
		filepicker:            fp,
		resultsCache:          make(map[string]string),
		noConstructsMsg:       "No constructs found",
		supportedFileTypes:    supportedFT,
		unsupportedFileMsg:    unsupportedFTMsg.String(),
		fileExplorerPaneWidth: fpWidth,
		showHelp:              true,
	}

	return m
}

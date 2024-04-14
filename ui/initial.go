package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/dstll/filepicker"
)

func InitialModel() model {

	fp := filepicker.New()
	supportedFT := []string{".go", ".scala", ".py"}

	unsupportedFTMsg := "dstll will show constructs for the following file types:\n"
	for _, ft := range supportedFT {
		unsupportedFTMsg += fmt.Sprintf("%s\n", ft)
	}

	fpWidth := 40

	fp.AllowedTypes = supportedFT
	fp.AutoHeight = false
	fp.Width = fpWidth
	fp.Styles.Selected.Foreground(lipgloss.Color(ActiveHeaderColor))
	fp.Styles.Cursor.Foreground(lipgloss.Color(ActiveHeaderColor))
	fp.Styles.DisabledFile.Foreground(lipgloss.Color(DisabledFileColor))
	fp.Styles.Directory.Foreground(lipgloss.Color(DirectoryColor))

	m := model{
		filepicker:            fp,
		resultsCache:          make(map[string]string),
		noConstructsMsg:       "No constructs found",
		supportedFileTypes:    supportedFT,
		unsupportedFileMsg:    unsupportedFTMsg,
		fileExplorerPaneWidth: fpWidth,
	}

	return m
}

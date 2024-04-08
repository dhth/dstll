package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/dhth/layitout/filepicker"
	"github.com/dhth/layitout/tsutils"
)

func InitialModel() model {

	fp := filepicker.New()
	supportedFT := []string{".go", ".scala", ".py"}

	unsupportedFTMsg := "layitout will show constructs for the following file types:\n"
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
		resultsCache:          make(map[string]tsutils.Result),
		noConstructsMsg:       "No constructs found",
		supportedFileTypes:    supportedFT,
		unsupportedFileMsg:    unsupportedFTMsg,
		fileExplorerPaneWidth: fpWidth,
	}

	return m
}

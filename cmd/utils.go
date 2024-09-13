package cmd

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func expandTilde(path string, homeDir string) string {
	pathWithoutTilde, found := strings.CutPrefix(path, "~/")
	if !found {
		return path
	}
	return filepath.Join(homeDir, pathWithoutTilde)
}

func createDir(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		mkDirErr := os.MkdirAll(path, 0o755)
		if mkDirErr != nil {
			return mkDirErr
		}
	}
	return nil
}

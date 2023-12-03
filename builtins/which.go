package builtins

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FindCommand searches for a command in the PATH environment variable.
func FindCommand(w io.Writer, commandName string) error {
	pathEnv := os.Getenv("PATH")
	dirs := filepath.SplitList(pathEnv)

	for _, dir := range dirs {
		fullPath := filepath.Join(dir, commandName)
		if fileExists(fullPath) {
			_, err := fmt.Fprintln(w, fullPath)
			return err
		}
	}
	return fmt.Errorf("command not found: %s", commandName)
}

// fileExists checks if a file exists and is not a directory.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

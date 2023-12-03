package builtins

import (
	"fmt"
	"io"
	"os"
)

// WorkingDirectoryGetter is an interface for getting the current working directory.
type WorkingDirectoryGetter interface {
	Getwd() (dir string, err error)
}

// OsWorkingDirectoryGetter implements WorkingDirectoryGetter using the os package.
type OsWorkingDirectoryGetter struct{}

// Getwd returns the current working directory using os.Getwd.
func (OsWorkingDirectoryGetter) Getwd() (string, error) {
	return os.Getwd()
}

// PrintWorkingDirectory prints the current working directory to the provided writer.
func PrintWorkingDirectory(w io.Writer, wdGetter WorkingDirectoryGetter) error {
	wd, err := wdGetter.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, wd)
	return err
}

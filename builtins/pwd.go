package builtins

import (
	"fmt"
	"io"
	"os"
)

type WorkingDirectoryGetter interface {
	Getwd() (dir string, err error)
}

type OsWorkingDirectoryGetter struct{}

func (OsWorkingDirectoryGetter) Getwd() (string, error) {
	return os.Getwd()
}

func PrintWorkingDirectory(w io.Writer, wdGetter WorkingDirectoryGetter) error {
	wd, err := wdGetter.Getwd()
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, wd)
	return err
}

package builtins

import (
	"fmt"
	"io"
	"strings"
)

func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}

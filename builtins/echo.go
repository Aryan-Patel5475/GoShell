// echo.go
package builtins

import (
	"fmt"
	"io"
	"strings"
)

// Echo prints its arguments joined by spaces to the provided writer
func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}

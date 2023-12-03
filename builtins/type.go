// type.go
package builtins

import (
	"fmt"
	"io"
	"os/exec"
)

// BuiltIns should be a map containing all built-in commands your shell supports.
// It needs to be maintained with the actual built-ins.
var BuiltIns = map[string]bool{
	"cd":     true,
	"env":    true,
	"exit":   true,
	"echo":   true,
	"pwd":    true,
	"repeat": true,
	"which":  true,
}

func TypeCommand(w io.Writer, command string) error {
	if _, ok := BuiltIns[command]; ok {
		fmt.Fprintf(w, "%s is a built-in command\n", command)
		return nil
	}

	// Here you might want to check for aliases if your shell supports them.
	// If it's an alias, print that and return.

	_, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("type: %s: not found", command)
	}

	fmt.Fprintf(w, "%s is a file\n", command)
	return nil
}

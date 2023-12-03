package builtins

import (
	"fmt"
	"io"
	"os/exec"
)

var BuiltIns = map[string]bool{
	"cd":     true,
	"env":    true,
	"exit":   true,
	"echo":   true,
	"pwd":    true,
	"repeat": true,
	"which":  true,
	"type":   true,
}

func TypeCommand(w io.Writer, command string) error {
	if _, ok := BuiltIns[command]; ok {
		fmt.Fprintf(w, "%s is a built-in command\n", command)
		return nil
	}

	_, err := exec.LookPath(command)
	if err != nil {
		return fmt.Errorf("type: %s: not found", command)
	}

	fmt.Fprintf(w, "%s is a file\n", command)
	return nil
}

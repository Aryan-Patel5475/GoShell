package builtins

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

func RepeatCommand(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments for repeat command")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid repeat count: %s", args[0])
	}

	command := args[1]

	for i := 0; i < count; i++ {
		cmd := exec.Command(command, args[2:]...)
		cmd.Stdout = w
		cmd.Stderr = w

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error executing command: %v", err)
		}
	}

	return nil
}

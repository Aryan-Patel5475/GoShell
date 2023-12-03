package builtins

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

// RepeatCommand repeats the given command a specified number of times.
func RepeatCommand(w io.Writer, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments for repeat command")
	}

	// Parse the number of times to repeat the command.
	count, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid repeat count: %s", args[0])
	}

	// The rest of the arguments form the command to repeat.
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

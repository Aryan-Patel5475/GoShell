package main

import (
	"bytes"
	"io"
	"os/exec"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			exit := make(chan struct{}, 2)
			// run the loop for 10ms
			go runLoop(tt.args.r, w, errW, exit)
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}

func TestHandleInput(t *testing.T) {

	sysEnvOutput, err := exec.Command("env").CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to execute system env command: %v", err)
	}
	/*
		expectedPwd, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current working directory: %v", err)
		}
	*/

	testCases := []struct {
		name        string
		input       string
		expectedOut string
		expectedErr string
	}{
		{"Echo Command", "echo Hello", "Hello\n", ""},
		{"Invalid Command", "notacommand", "", "not a built-in command"},
		{"cd command", "cd", "", ""},
		{"Invalid Command", "notacommand", "", "not a built-in command"},
		{"env command", "env", string(sysEnvOutput), ""},
		//{"pwd command", "pwd", expectedPwd + "\n", ""},
		{"repeat commmand", "repeat 3 echo Hello", "Hello\nHello\nHello\n", ""},
		{"which command", "which echo", "/bin/echo\n", ""},
		{"type command", "type echo", "echo is a built-in command\n", ""},
		// Add more cases for each command
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			exit := make(chan struct{}, 1)

			// Call handleInput
			err := handleInput(output, tc.input, exit)

			// Check for expected errors
			if err != nil {
				if tc.expectedErr == "" {
					t.Errorf("Unexpected error: %s", err)
				} else if !strings.Contains(err.Error(), tc.expectedErr) {
					t.Errorf("Expected error '%s', got '%s'", tc.expectedErr, err.Error())
				}
			} else {
				if tc.expectedErr != "" {
					t.Errorf("Expected error '%s' but got none", tc.expectedErr)
				}
			}

			// Check the output
			if output.String() != tc.expectedOut {
				t.Errorf("Expected output '%s', got '%s'", tc.expectedOut, output.String())
			}
		})
	}
}

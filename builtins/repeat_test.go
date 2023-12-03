package builtins_test

import (
	"bytes"
	"github.com/Aryan-Patel5475/Go-Shell/builtins"
	"testing"
)

func TestRepeatCommand(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantOutput string
		wantErr    bool
	}{
		{
			name:       "Repeat Echo 3 Times",
			args:       []string{"3", "echo", "hello"},
			wantOutput: "hello\nhello\nhello\n",
			wantErr:    false,
		},
		{
			name:    "Invalid Repeat Count",
			args:    []string{"invalid", "echo", "hello"},
			wantErr: true,
		},
		{
			name:    "No Command Provided",
			args:    []string{"3"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := builtins.RepeatCommand(&buf, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RepeatCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && buf.String() != tt.wantOutput {
				t.Errorf("RepeatCommand() got output = %v, want %v", buf.String(), tt.wantOutput)
			}
		})
	}
}

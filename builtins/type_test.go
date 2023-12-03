// type_test.go
package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Aryan-Patel5475/Go-Shell/builtins"
)

func TestTypeCommand(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
		wantErr bool
	}{
		{"built-in command", "cd", "cd is a built-in command\n", false},
		{"external command", "ls", "ls is a file\n", false},
		{"invalid command", "nonexistentcommand", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := builtins.TypeCommand(&buf, tt.command)

			if (err != nil) != tt.wantErr {
				t.Errorf("TypeCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got := buf.String(); got != tt.want {
				t.Errorf("TypeCommand() got = %v, want %v", got, tt.want)
			}
		})
	}
}

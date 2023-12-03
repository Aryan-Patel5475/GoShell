// echo_test.go
package builtins_test

import (
	"bytes"
	"testing"

	"github.com/Aryan-Patel5475/Go-Shell/builtins"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "single word",
			args:     []string{"hello"},
			expected: "hello\n",
		},
		{
			name:     "multiple words",
			args:     []string{"hello", "world"},
			expected: "hello world\n",
		},
		{
			name:     "no arguments",
			args:     []string{},
			expected: "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buffer bytes.Buffer

			err := builtins.Echo(&buffer, tt.args...)
			if err != nil {
				t.Errorf("Echo() error = %v", err)
			}
			if got := buffer.String(); got != tt.expected {
				t.Errorf("Echo() = %v, want %v", got, tt.expected)
			}
		})
	}
}

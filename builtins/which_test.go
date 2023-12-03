package builtins_test

import (
	"bytes"
	"github.com/Aryan-Patel5475/Go-Shell/builtins"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFindCommand(t *testing.T) {
	tempDir := t.TempDir()
	os.Setenv("PATH", tempDir)

	mockCommand := filepath.Join(tempDir, "mockCommand")
	if err := os.WriteFile(mockCommand, []byte("#!/bin/sh\n"), 0755); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name       string
		command    string
		wantErr    bool
		wantOutput string
	}{
		{
			name:       "Existing Command",
			command:    "mockCommand",
			wantErr:    false,
			wantOutput: mockCommand,
		},
		{
			name:    "Non-existing Command",
			command: "nonExistingCommand",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := builtins.FindCommand(&buf, tt.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && strings.TrimSpace(buf.String()) != tt.wantOutput {
				t.Errorf("FindCommand() = %v, want %v", buf.String(), tt.wantOutput)
			}
		})
	}
}

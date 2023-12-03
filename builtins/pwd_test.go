package builtins_test

import (
	"bytes"
	"fmt"
	"github.com/Aryan-Patel5475/Go-Shell/builtins"
	"os"
	"testing"
)

// MockFailWorkingDirectoryGetter is a mock that simulates failure in getting the working directory.
type MockFailWorkingDirectoryGetter struct{}

// Getwd simulates a failure in getting the working directory.
func (MockFailWorkingDirectoryGetter) Getwd() (string, error) {
	return "", fmt.Errorf("mock error")
}

func TestPrintWorkingDirectory(t *testing.T) {
	var buf bytes.Buffer
	err := builtins.PrintWorkingDirectory(&buf, builtins.OsWorkingDirectoryGetter{})
	if err != nil {
		t.Fatalf("PrintWorkingDirectory() error: %v", err)
	}

	got := buf.String()
	want, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd() error: %v", err)
	}

	if got != want+"\n" {
		t.Errorf("PrintWorkingDirectory() = %v, want %v", got, want)
	}
}

func TestPrintWorkingDirectoryFail(t *testing.T) {
	var buf bytes.Buffer
	err := builtins.PrintWorkingDirectory(&buf, MockFailWorkingDirectoryGetter{})
	if err == nil {
		t.Fatalf("Expected error from PrintWorkingDirectory(), got nil")
	}
}

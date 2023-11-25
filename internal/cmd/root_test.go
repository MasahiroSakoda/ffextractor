// Package cmd provides command using cobra cli library
package cmd

import (
	"path/filepath"
	"runtime"

	"testing"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	rootDir = filepath.Join(filepath.Dir(b), "../..")
)

func TestExecute(t *testing.T) {
	if Execute(nil) != 0 {
		t.Errorf("exit code must be 0")
	}
}

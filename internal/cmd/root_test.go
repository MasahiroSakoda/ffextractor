package cmd

import (
	"testing"
)

func TestExecute(t *testing.T) {
	if Execute(nil) != 0 {
		t.Errorf("exit code must be 0")
	}
}

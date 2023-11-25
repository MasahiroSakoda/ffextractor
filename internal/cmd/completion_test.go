// Package cmd provides command using cobra cli library
package cmd

import (
	"bytes"
	"io"
	"testing"
)

func TestCompletionCmd(t *testing.T) {
	tests := []struct{
		name    string
		param   string
		output  string
		wantErr bool
	}{
		{ param: "bash", output: "bash completion for bash",           wantErr: false, name: "return bash completion " },
		{ param: "zsh",  output: "zsh completion for zsh",             wantErr: false, name: "return zsh completion" },
		{ param: "fish", output: "fish completion for fish",           wantErr: false, name: "return fish completion" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newCompletionCmd()
			cmd.SetArgs([]string{tt.param})
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			err := cmd.Execute()
			if err != nil {
				t.Fatalf("newCompletionCmd() execute error: %v", err)
			}
			o, err := io.ReadAll(b)
			if err != nil {
				t.Fatalf("stdout capturing error: %v", err)
			}
			var output = string(o)
			if output == "" {
				t.Errorf("Unexpected output: %s", output)
			}
		})
	}
}

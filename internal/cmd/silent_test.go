package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestSilentCmd(t *testing.T) {
	tests := []struct{
		name   string
		cmd    string
		param  string
		output string
		expect bool
	}{
		{
			cmd: "silent",
			param: "",
			output: "Usage:",
			expect: true,
			name: "return exit(1) without parameter",
		},
		{
			cmd: "silent",
			param: "fail.txt",
			output: "Usage:",
			expect: true,
			name: "return exit(1) with wrong parameter",
		},
		{
			cmd: "silent",
			param: "invalid.mp3",
			output: "Usage:",
			expect: true,
			name: "return exit(1) with wrong parameter",
		},
		{
			cmd: "silent",
			param: rootDir + "/testdata/ffmpeg/sine.mp3",
			output: "Usage:",
			expect: false,
			name: "return exit(0) with with correct parameter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{constants.CommandName, tt.cmd, tt.param}
			cmd := newSilentCmd()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.Execute()
			out, err := io.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Contains(t, string(out), tt.output)
		})
	}
}

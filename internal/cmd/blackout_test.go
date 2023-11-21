package cmd

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestBlackoutCmd(t *testing.T) {
	tests := []struct{
		name   string
		cmd    string
		param  string
		output string
		wantErr bool
	}{
		{
			cmd: "blackout",
			param: "",
			output: "Usage:",
			wantErr: true,
			name: "return exit(1) without parameter",
		},
		{
			cmd: "blackout",
			param: "fail.txt",
			output: "Usage:",
			wantErr: true,
			name: "return exit(1) with wrong parameter",
		},
		{
			cmd: "blackout",
			param: "invalid.mp3",
			output: "Usage:",
			wantErr: true,
			name: "return exit(1) with wrong parameter",
		},
		{
			cmd: "blackout",
			param: rootDir + "/testdata/ffmpeg/sine.mp3",
			output: "Usage:",
			wantErr: false,
			name: "return exit(0) with with correct parameter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = []string{constants.CommandName, tt.cmd, tt.param}
			cmd := newBlackoutCmd()
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				if tt.wantErr {
					t.Errorf("cmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			out, err := io.ReadAll(b)
			if err != nil {
				t.Fatal(err)
			}
			assert.Contains(t, string(out), tt.output)
		})
	}
}

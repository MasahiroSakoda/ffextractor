package cmd

import (
	"testing"
)

func TestSilentCmd(t *testing.T) {
	tests := []struct{
		name   string
		cmd    string
		param  string
		wantErr bool
	}{
		{
			cmd: "silent",
			param: "",
			wantErr: true,
			name: "return exit(1) without parameter",
		},
		{
			cmd: "silent",
			param: rootDir + "fail.txt",
			wantErr: true,
			name: "return invalid parameter with file does not exist",
		},
		{
			cmd: "silent",
			param: "invalid.mp3",
			wantErr: true,
			name: "return invalid parameter with media file does not exist",
		},
		{
			cmd: "silent",
			param: rootDir + "/testdata/ffmpeg/sine.mp3",
			wantErr: false,
			name: "not detected with with noisy file",
		},
		{
			cmd: "silent",
			param: rootDir + "/testdata/ffmpeg/mixed.mp3",
			wantErr: false,
			name: "return exit(0) with with correct parameter",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newSilentCmd()
			cmd.SetArgs([]string{tt.param})
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				if tt.wantErr {
					t.Errorf("cmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
				t.Logf("parameter: %s", tt.param)
			}
		})
	}
}

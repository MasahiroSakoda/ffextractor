// Package cmd provides command using cobra cli library
package cmd

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigCmd(t *testing.T) {
	tests := []struct{
		name    string
		key     string
		value   string
		wantErr bool
	}{
		{ key: "overwrite",  value: "true",  wantErr: false, name: "Successfully configured with true" },
		{ key: "overwrite",  value: "false", wantErr: false, name: "Successfully configured with false" },
		{ key: "overwrite",  value: "fail",  wantErr: true,  name: "Failed with string value" },
		{ key: "overwrite",  value: "123",   wantErr: true,  name: "Failed with numbered value" },
		{ key: "annotation", value: "_A",    wantErr: false, name: "Successfully configure with string value" },
		{ key: "annotation", value: "",      wantErr: true,  name: "Failed without parameter" },
		{ key: "threshold",  value: "15",    wantErr: false, name: "Successfully configured with int value" },
		{ key: "threshold", value: "-5",     wantErr: true,  name: "Successfully configured with negative value" },
		{ key: "threshold", value: "abc",    wantErr: false, name: "Failed with string value" },

		{ key: "silence_duration",  value: "ABC", wantErr: true,  name: "Failed with string value" },
		{ key: "silence_duration",  value: "---", wantErr: true,  name: "Failed with symbol" },
		{ key: "silence_duration",  value: "60",  wantErr: false, name: "Successfully configured numbered value" },
		{ key: "blackout_duration", value: "DEF", wantErr: true,  name: "Failed with string value" },
		{ key: "blackout_duration", value: "^^^", wantErr: true,  name: "Failed with symbol value" },
		{ key: "blackout_duration", value: "30",  wantErr: false, name: "Successfully configured numbered value" },

		{ key: "split_with_encode",  value: "true",  wantErr: false, name: "Successfully configured with true" },
		{ key: "split_with_encode",  value: "false", wantErr: false, name: "Successfully configured with false" },
		{ key: "split_with_encode",  value: "fail",  wantErr: true,  name: "Failed with string value" },
		{ key: "concat_with_encode", value: "true",  wantErr: false, name: "Successfully configured with true" },
		{ key: "concat_with_encode", value: "false", wantErr: false, name: "Successfully configured with false" },
		{ key: "concat_with_encode", value: "fail",  wantErr: true,  name: "Failed with string value" },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := newConfigCmd()
			cmd.SetArgs([]string{tt.key, tt.value})
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			err := cmd.Execute()
			if (err != nil) != tt.wantErr {
				if tt.wantErr {
					t.Errorf("cmd.Execute() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			o, err := io.ReadAll(b)
			if err != nil {
				t.Errorf("io.ReadAll() error: %v", err)
			}
			var output = string(o)
			if (err != nil) != tt.wantErr {
				assert.Contains(t, output, "Usage")
			}
		})
	}
}

// Package extractor provides detect -> split -> concat function using `ffmpeg`
package extractor

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDataPath(filename string) string {
	absPath, _ := filepath.Abs("../../testdata/ffmpeg/" + filename)
	return filepath.Clean(absPath)
}

func TestDetectSilentSegments(t *testing.T) {
	tests := []struct {
		file    string
		name    string
		wantErr bool
	}{
		{ file: testDataPath("invalid.mp3"), wantErr: true,  name: "Return error with invalid file" },
		{ file: testDataPath("sine.mp3"),    wantErr: true,  name: "Undetected with noisy file" },
		{ file: testDataPath("silence.mp3"), wantErr: true,  name: "Undetected with silent file" },
		{ file: testDataPath("mixed.mp3"),   wantErr: false, name: "Detected silence parts with proper file" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectSilentSegments(testDataPath(tt.file))
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.wantErr, got != nil)
			}
		})
	}
}

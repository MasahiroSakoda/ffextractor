package ffmpeg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testDataPath(filename string) string {
	return "./testdata/ffmpeg/" + filename
}

func TestDetectSilence(t *testing.T) {
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
			got, err := DetectSilence(testDataPath(tt.file))
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.wantErr, got != nil)
			}
		})
	}
}

func TestDetectBlackout(t *testing.T) {
	tests := []struct {
		file    string
		name    string
		wantErr bool
	}{
		// TODO: Uncomment after implement DetectBlackout()
		// { file: testDataPath("invalid.mp4"),  wantErr: true, name: "Return error with invalid file"},
		// { file: testDataPath("blackout.mp4"), wantErr: true, name: "Undetected with blackout only video"},
		// { file: testDataPath("whiteout.mp4"), wantErr: true, name: "Undetected with whiteout only video"},
		// { file: testDataPath("mixed.mp4"),    wantErr: false, name: "Detected blackout parts with proper file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectBlackout(testDataPath(tt.file))
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.wantErr, got != nil)
			}
		})
	}
}

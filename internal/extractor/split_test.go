// Package extractor provides detect -> split -> concat function using `ffmpeg`
package extractor

import (
	"os"
	"runtime"
	"path/filepath"
	"testing"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	rootDir = filepath.Join(filepath.Dir(b), "../..")
)

func TestSplitDetectedSegment(t *testing.T) {
	tests := []struct{
		name   string
		path   string
		encode bool
		expect bool
	}{
		{
			path: testDataPath("mixed.mp3"),
			encode: false,
			expect: true,
			name: "split detected segments without encode",
		},
		{
			path: testDataPath("mixed.mp3"),
			encode: true,
			expect: true,
			name: "split detected segments with encode",
		},
	}

	for _, tt := range tests {
		// create temp directory
		dirPrefix := constants.CommandName + "_" + util.GetFilenameFromPath(tt.path)
		tempDir, err := os.MkdirTemp("", dirPrefix)
		if err != nil {
			t.Errorf("%s: %v", constants.ErrMkdir, err)
		}
		defer func() {
			os.RemoveAll(tempDir)
		}()

		t.Run(tt.name, func(t *testing.T) {
			segments, err := DetectSilentSegments(tt.path)
			if err != nil {
				t.Errorf("%s: %v", constants.ErrSilenceDetect, err)
			}

			for _, segment := range segments {
				err := SplitDetectedSegment(segment, tempDir)
				if err != nil {
					t.Errorf("%s: %v", constants.ErrSplitSegment, err)
				}
			}
		})
	}
}

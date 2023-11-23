package extractor

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/MasahiroSakoda/ffextractor/internal/config"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"

	// ffmeg "github.com/u2takey/ffmpeg-go"
)

// SplitDetectedSegment returns error to extract media segment exclude detected parts
func SplitDetectedSegment(segment segment.Model, tempDir string) error {
	output := filepath.Join(tempDir, segment.Output)
	encode := config.Root.Encode.SplitWithEncode

	args := []string {
		"-y",
		"-ss", strconv.FormatFloat(segment.Start, 'f', -1, 64),
		"-i",  segment.Input,
		"-t",  strconv.FormatFloat(segment.Duration, 'f', -1, 64),
	}
	if encode {
		args = append(args, "-c", "copy")
	}
	args = append(args, output)
	fmt.Println(args)
	// output, err := exec.Command("ffmpeg", args...).CombinedOutput()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(string(output))
	return nil
}

package extractor

import (
	"path/filepath"

	"github.com/MasahiroSakoda/ffextractor/internal/config"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"

	fg "github.com/u2takey/ffmpeg-go"
)

// SplitDetectedSegment returns error to extract media segment exclude detected parts
func SplitDetectedSegment(segment segment.Model, tempDir string) error {
	output := filepath.Join(tempDir, segment.Output)
	encode := config.Root.Encode.SplitWithEncode
	// fmt.Println(output)

	a, err := fg.Probe(segment.Output)
	if err != nil { return err }

	totalDuration, err := probeDuration(a)
	if err != nil { return err }

	var outArgs fg.KwArgs
	if !encode {
		outArgs = fg.KwArgs{"t": segment.Duration}
	} else {
		outArgs = fg.KwArgs{"t": segment.Duration, "c": "copy"}
	}

	err = fg.Input(segment.Input, fg.KwArgs{"ss": segment.Start}).
			Output(output, outArgs).
			OverWriteOutput().
			GlobalArgs("-progress", "unix://" + TempSock(totalDuration)).
			Compile().
			Run()
	return err
}

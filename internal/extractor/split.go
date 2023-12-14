// Package extractor provides detect -> split -> concat function using `ffmpeg`
package extractor

import (
	"fmt"
	"path/filepath"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/config"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"

	fg "github.com/u2takey/ffmpeg-go"
)

// SplitDetectedSegment returns error to extract media segment exclude detected parts
func SplitDetectedSegment(segment segment.Model, tempDir string) error {
	output := filepath.Join(tempDir, segment.Output)
	c, err := config.Load()
	if err != nil { return err }
	encode := c.Encode.SplitWithEncode

	var inArgs  fg.KwArgs = fg.KwArgs{"ss": segment.Start}
	var outArgs fg.KwArgs
	if !encode {
		outArgs = fg.KwArgs{"t": segment.Duration}
	} else {
		outArgs = fg.KwArgs{"t": segment.Duration, "c": "copy"}
	}

	// a, err := fg.Probe(segment.Output)
	// if err != nil { return err }
	// totalDuration, err := probeDuration(a)
	// if err != nil { return err }

	err = fg.Input(segment.Input, inArgs).
			Output(output, outArgs).
			Silent(true).
			OverWriteOutput().
			// TODO: add progress function
			// GlobalArgs("-progress", "unix://" + TempSock(totalDuration)).
			// Compile().
			Run()
	return fmt.Errorf("%s: %+v", constants.ErrSplitSegment, err)
}

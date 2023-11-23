package extractor

import (
	"os/exec"
	"fmt"
	"path/filepath"
	"strings"
	"strconv"

	"github.com/MasahiroSakoda/ffextractor/internal/config"
	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"
	"github.com/MasahiroSakoda/ffextractor/internal/util"
)

func DetectSilence(src string) ([]segment.Model, error) {
	extract   := config.Root.Extract
	threshold := extract.Threshold
	duration  := extract.SilenceDuration

	args := []string{
		"-i", src,
		"-af", fmt.Sprintf("silencedetect=noise=%ddB:d=%2.3f", threshold, duration),
		"-f", "null", "-"}
	output, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		return []segment.Model{}, err
	}
	return parseDetectedSegments(src, output)
}

func parseDetectedSegments(src string, data []byte) (segments []segment.Model, err error) {
	var s segment.Model
	s.Start = 0
	correction   := 0.0
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, "silence_end") {
			seconds, err := util.ConvertToSeconds(util.GetStringInBetween(line, "silence_end: ", " "))
			if err == nil {
				s.End = seconds + correction
				s.Input = src
				s.Duration = s.End - s.Start
				if s.Duration > 0.25 {
					segments = append(segments, s)
				}
				s.Start  = seconds + correction
				s.Output = fmt.Sprintf("%s_%s%s",
					constants.CommandName,
					strconv.FormatFloat(s.Start, 'f', -1, 64),
					filepath.Ext(s.Input),
				)
			} else {
				fmt.Printf("%s", err)
			}
		} else if strings.Contains(line, "time=") {
			seconds, err := util.ConvertToSeconds(util.GetStringInBetween(line, "time=", " "))
			if err == nil {
				s.End = seconds
				s.Duration = s.End - s.Start
				s.Input = src
				if s.Duration < 0.25 && len(segments) > 0 {
					segments[len(segments)-1].End = seconds
					segments[len(segments)-1].Duration = segments[len(segments)-1].End - segments[len(segments)-1].Start
				} else {
					segments = append(segments, s)

				}
			} else {
				fmt.Printf("%s", err)
			}
		}
	}
	newSegments := make([]segment.Model, len(segments))
	i := 0

	for _, seg := range segments {
		if seg.Duration > 0.1 {
			newSegments[i] = seg
			i++
		}
	}
	if i == 0 {
		err = fmt.Errorf("could not find any segments")
		return
	}

	newSegments = newSegments[:i]
	return newSegments, nil
}

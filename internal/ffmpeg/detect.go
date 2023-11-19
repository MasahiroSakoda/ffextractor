package ffmpeg

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/MasahiroSakoda/ffextractor/internal/util"
	"github.com/MasahiroSakoda/ffextractor/internal/model"
	"github.com/MasahiroSakoda/ffextractor/internal/config"
)

// DetectSilence returns stdout for detected silence parts
func DetectSilence(src string) (segments []model.AudioSegment, err error) {
	extract := config.Root.Extract
	var threshold = extract.Threshold
	var duration  = extract.SilenceDuration
	correction   := 0.0
	args := []string{
		"-i", src,
		"-af", fmt.Sprintf("silencedetect=noise=%ddB:d=%2.3f", threshold, duration),
		"-f", "null", "-"}
	output, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}

	var segment model.AudioSegment
	segment.Start = 0
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, "silence_end") {
			seconds, err := util.ConvertToSeconds(util.GetStringInBetween(line, "silence_end: ", " "))
			if err == nil {
				segment.End = seconds + correction
				segment.Filename = src
				segment.Duration = segment.End - segment.Start
				if segment.Duration > 0.25 {
					segments = append(segments, segment)
				}
				segment.Start = seconds + correction
			} else {
				fmt.Printf("%s", err)
			}
		} else if strings.Contains(line, "time=") {
			seconds, err := util.ConvertToSeconds(util.GetStringInBetween(line, "time=", " "))
			if err == nil {
				segment.End = seconds
				segment.Duration = segment.End - segment.Start
				segment.Filename = src
				if segment.Duration < 0.25 && len(segments) > 0 {
					segments[len(segments)-1].End = seconds
					segments[len(segments)-1].Duration = segments[len(segments)-1].End - segments[len(segments)-1].Start
				} else {
					segments = append(segments, segment)

				}
			} else {
				fmt.Printf("%s", err)
			}
		}
	}
	newSegments := make([]model.AudioSegment, len(segments))
	i := 0

	for _, segment := range segments {
		if segment.Duration > 0.1 {
			newSegments[i] = segment
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

// DetectBlackout returns stdout for detected blackout parts
func DetectBlackout(src string) (segments []model.VideoSegment, err error) {
	duration  := config.Root.Extract.BlackoutDuration
	args := []string{
		"-i", src,
		"-af", fmt.Sprintf("blackdetect=%2.3f", duration),
		"-f", "null", "-"}
	output, err := exec.Command("ffmpeg", args...).CombinedOutput()
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", string(output))

	var segment model.VideoSegment
	segment.Start = 0

	// TODO: collect detected parts
	// for _, line := range strings.Split(string(output), "\n") {
	// 	if strings.Contains(line, "") {
	// 	} else if strings.Contains(line, "time=") {
	// 	}
	// }

	return segments, nil
}

package ffmpeg

import (
	"fmt"
	"os/exec"

	"github.com/MasahiroSakoda/ffextractor/internal/config"
)

// DetectSilence returns stdout for detected silence parts
func DetectSilence(src string) ([]byte, error) {
	threshold := config.Root.Extract.Threshold
	duration  := config.Root.Extract.SilenceDuration
	args := []string{
		"-i", src,
		"-af", fmt.Sprintf("silencedetect=noise=%ddB:d=%2.3f", threshold, duration),
		"-f", "null", "-"}
	return exec.Command("ffmpeg", args...).CombinedOutput()
}

package ffmpeg

import (
	"os/exec"
)

func IsInstalled() bool {
	args := []string{"--help"}
	_, err := exec.Command("ffmpeg", args...).CombinedOutput()
	return err == nil
}

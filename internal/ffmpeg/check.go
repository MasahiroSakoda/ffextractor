// Package ffmpeg provides function related `ffmpeg`
package ffmpeg

import (
	"os/exec"
)

// IsInstalled returns whether `ffmpeg` is installed
func IsInstalled() bool {
	args := []string{"--help"}
	_, err := exec.Command("ffmpeg", args...).CombinedOutput()
	return err == nil
}

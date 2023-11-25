package main

import (
	"os"

	"github.com/MasahiroSakoda/ffextractor/internal/cmd"
	"github.com/MasahiroSakoda/ffextractor/internal/ffmpeg"

	"github.com/sirupsen/logrus"
)

func main() {
	if !ffmpeg.IsInstalled() {
		logrus.Errorf("ffmpeg is not installed")
		os.Exit(1)
	}

	os.Exit(cmd.Execute(os.Args[1:]))
}

package main

import (
	"os"
	"fmt"
	"github.com/MasahiroSakoda/ffextractor/internal/cmd"
	"github.com/MasahiroSakoda/ffextractor/internal/ffmpeg"
)

func main() {
	if !ffmpeg.IsInstalled() {
		fmt.Println("ffmpeg is not installed")
		os.Exit(1)
	}
	os.Exit(cmd.Execute(os.Args[1:]))
}

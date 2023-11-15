package main

import (
	"github.com/MasahiroSakoda/ffextractor/internal/cmd"
)

var (
	version string
	commit  string
	date    string
	buildBy string
)

func main() {
	cmd.Execute()
}

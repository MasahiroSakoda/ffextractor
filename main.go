package main

import (
	"os"
	"github.com/MasahiroSakoda/ffextractor/internal/cmd"
)

func main() {
	os.Exit(cmd.Execute(os.Args[1:]))
}

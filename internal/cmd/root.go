package cmd

import (
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string

	flagConfig string
)

var rootCmd = &cobra.Command{
	Use:   "ffextractor",
	Short: "Automates terminal operations",
	Long:  "Automates terminal operations.",
}

// Execute : root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Init : initialize at first launch
func Init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version
	// _ = notifyNewRelease(os.Stderr)

	rootCmd.AddCommand(
		silentCmd,
		blackoutCmd,
		configCmd,
	)

	for _, cmd := range []*cobra.Command{
		silentCmd,
		blackoutCmd,
		configCmd,
	} {
		// --config
		cmd.Flags().StringVarP(&flagConfig, "config", "c", "./clive.yml", "config file name")
	}
}

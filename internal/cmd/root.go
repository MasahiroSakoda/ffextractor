// Package cmd provides command using cobra cli library
package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string
)

var subCmds = []string{
	"silent",
	"blackout",
	"config",
	"completion",
}

var rootCmd = &cobra.Command{
	Use:   "ffextractor",
	Short: "Automates terminal operations",
	Long:  "Automates terminal operations.",
	ValidArgs: subCmds,
	Args: cobra.MinimumNArgs(3),
}

// Execute : root command
func Execute(args []string) int {
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(rootCmd.OutOrStderr(), "Err: %s\n", err)
		return 1
	}
	return 0
}

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version
	_ = notifyNewRelease(os.Stderr)

	for _, cmd := range []*cobra.Command{
		newSilentCmd(),
		newBlackoutCmd(),
		newConfigCmd(),
		newCompletionCmd(),
	} {
		rootCmd.AddCommand(cmd)
	}
}

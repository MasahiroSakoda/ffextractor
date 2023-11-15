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

var rootCmd = &cobra.Command{
	Use:   "ffextractor",
	Short: "Automates terminal operations",
	Long:  "Automates terminal operations.",
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
		silentCmd,
		blackoutCmd,
		configCmd,
	} {
		rootCmd.AddCommand(cmd)
	}
}

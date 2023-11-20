package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/MasahiroSakoda/ffextractor/internal/config"

	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

var (
	version string
)

var subCmds = []string{
	"silent",
	"blackout",
	"config",
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

	_, err := config.Root.Load("")
	if err != nil {
		logrus.Fatalf("cli: failed to load config")
	}

	configCmd.AddCommand(
		overwriteCmd,
		annotationCmd,
		thresholdCmd,
		silenceDurationCmd,
		blackoutDurationCmd,
	)

	completionCmd.AddCommand(
		completionBashCmd,
		completionZshCmd,
		completionFishCmd,
	)

	for _, cmd := range []*cobra.Command{
		silentCmd,
		blackoutCmd,
		completionCmd,
		configCmd,
	} {
		rootCmd.AddCommand(cmd)
	}
}

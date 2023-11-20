package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"
	"github.com/sirupsen/logrus"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure " + constants.CommandName + " options.",
	Long:  "Configure " + constants.CommandName + " options.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var overwriteCmd = &cobra.Command{
	Use:   "overwrite",
	Short: "Overwrite existing file",
	Long:  "Overwrite existing file (default: false)",
	Args:  cobra.ExactArgs(1),
	RunE:  runOverwriteCmd,
}

var annotationCmd = &cobra.Command{
	Use:   "annotation",
	Short: "Configure file suffix",
	Long:  "Configure file suffix (default: \"_merged\")",
	Args:  cobra.ExactArgs(1),
	RunE:  runAnnotationCmd,
}

var thresholdCmd = &cobra.Command{
	Use:   "threshold",
	Short: "Volume threshold to detect silence",
	Long:  "Volume threshold to detect silence (default: 50)[dB]",
	Args:  cobra.ExactArgs(1),
	RunE:  runThresholdCmd,
}

var silenceDurationCmd = &cobra.Command{
	Use:   "silence_duration",
	Short: "Duration to detect silence",
	Long:  "Duration to detect silence (default: 5.0)[sec]",
	Args:  cobra.ExactArgs(1),
	RunE:  runSilenceDurationCmd,
}

var blackoutDurationCmd = &cobra.Command{
	Use:   "blackout_duration",
	Short: "Duration to detect blackout",
	Long:  "Duration to detect blackout (default: 5.0)[sec]",
	Args:  cobra.ExactArgs(1),
	RunE:  runBlackoutDurationCmd,
}

func runConfigCmd(cmd *cobra.Command, _ []string) error {
	cmd.AddCommand(
		overwriteCmd,
		annotationCmd,
		thresholdCmd,
		silenceDurationCmd,
		blackoutDurationCmd,
	)
	return nil
}

func runOverwriteCmd(_ *cobra.Command, args []string) error {
	if !util.IsBoolean([]byte(args[0])) {
		logrus.Errorf("overwrite should use boolean value")
		os.Exit(1)
	}
	// TODO: save annotation config
	return nil
}

func runAnnotationCmd(_ *cobra.Command, args []string) error {
	if len(args[0]) > 0 {
		// TODO: save annotation config
	}
	return nil
}

func runThresholdCmd(_ *cobra.Command, args []string) error {
	if !util.IsInteger([]byte(args[0])) {
		logrus.Errorf("threshold should use non-negative integer value")
		os.Exit(1)
	}
	// TODO: save threshold config
	return nil
}

func runSilenceDurationCmd(_ *cobra.Command, args []string) error {
	if !util.IsInteger([]byte(args[0])) && !util.IsFloat([]byte(args[0])) {
		logrus.Errorf("silence_duration should use non-negative float value")
		os.Exit(1)
	}
	// TODO: save silence_duration config
	return nil
}

func runBlackoutDurationCmd(_ *cobra.Command, args []string) error {
	if !util.IsInteger([]byte(args[0])) && !util.IsFloat([]byte(args[0])) {
		logrus.Errorf("blackout_duration should use non-negative float value")
		os.Exit(1)
	}
	// TODO: save blackout_duration config
	return nil
}

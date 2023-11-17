package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/MasahiroSakoda/ffextractor/internal/util"
	"github.com/sirupsen/logrus"
)

func init() {
	rootCmd.AddCommand(newConfigCmd())
}

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command {
		Use:   "config",
		Short: "Configure ffextractor options.",
		Long:  "Configure ffextractor options.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(
		newOverwriteCmd(),
		newAnnotationCmd(),
		newThresholdCmd(),
		newSilenceDurationCmd(),
		newBlackoutDurationCmd(),
	)

	return cmd
}

func newOverwriteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "overwrite",
		Short: "Overwrite existing file",
		Long:  "Overwrite existing file (default: false)",
		Args:  cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !util.IsBoolean([]byte(args[0])) {
				logrus.Errorf("overwrite should use boolean value")
				os.Exit(1)
			}
			return nil
		},
	}
	return cmd
}

func newAnnotationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "annotation",
		Short: "Configure file suffix",
		Long:  "Configure file suffix (default: \"_merged\")",
		Args:  cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return cmd
}

func newThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "threshold",
		Short: "Volume threshold to detect silence",
		Long:  "Volume threshold to detect silence (default: 50)[dB]",
		Args:  cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !util.IsInteger([]byte(args[0])) {
				logrus.Errorf("threshold should use non-negative integer value")
				os.Exit(1)
			}
			return nil
		},
	}
	return cmd
}

func newSilenceDurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "silence_duration",
		Short: "Duration to detect silence",
		Long:  "Duration to detect silence (default: 5.0)[sec]",
		Args:  cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !util.IsInteger([]byte(args[0])) && !util.IsFloat([]byte(args[0])) {
				logrus.Errorf("silence_duration should use non-negative float value")
				os.Exit(1)
			}
			return nil
		},
	}
	return cmd
}

func newBlackoutDurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blackout_duration",
		Short: "Duration to detect blackout",
		Long:  "Duration to detect blackout (default: 5.0)[sec]",
		Args:  cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveDefault
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if !util.IsInteger([]byte(args[0])) && !util.IsFloat([]byte(args[0])) {
				logrus.Errorf("blackout_duration should use non-negative float value")
				os.Exit(1)
			}
			return nil
		},
	}
	return cmd
}

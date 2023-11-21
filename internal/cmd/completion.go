package cmd

import (
	"github.com/spf13/cobra"
)

func newCompletionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generates shell completion scripts",
		Long:  "Generates shell completion scripts",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		ValidArgs: []string{"bash", "zsh", "fish"},
		RunE:  func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				_ = cmd.GenBashCompletion(cmd.OutOrStderr())
			case "zsh":
				_ = cmd.GenZshCompletion(cmd.OutOrStderr())
			case "fish":
				_ = cmd.GenFishCompletion(cmd.OutOrStderr(), true)
			}
			return nil
		},
	}
}

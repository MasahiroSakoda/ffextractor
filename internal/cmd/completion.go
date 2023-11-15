package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCompletionCmd())
}

func newCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates shell completion scripts",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	return cmd
}

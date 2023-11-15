package cmd

import (
	"os"
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
	cmd.AddCommand(
		newCompletionBashCmd(),
		newCompletionZshCmd(),
	)
	return cmd
}

func newCompletionBashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenZshCompletion(os.Stdout)
		},
	}
	return cmd
}

func newCompletionZshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return cmd
}

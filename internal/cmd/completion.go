package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates shell completion scripts",
	Long:  "Generates shell completion scripts",
	Args:  cobra.MinimumNArgs(1),
	RunE:  runCompletionCmd,
}

var completionBashCmd = &cobra.Command {
	Use:   "bash",
	Short: "Generates bash completion scripts",
	Long:  "Generates bash completion scripts",
	RunE:  runCompletionBashCmd,
}

var completionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generates zsh completion scripts",
	Long:  "Generates zsh completion scripts",
	RunE:   runCompletionZshCmd,
}

var completionFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generates fish completion scripts",
	Long:  "Generates fish completion scripts",
	RunE:  runCompletionFishCmd,
}

func runCompletionCmd(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}

func runCompletionBashCmd(cmd *cobra.Command, _ []string) error {
	err := cmd.GenBashCompletion(os.Stdout)
	if err != nil {
		os.Exit(1)
		return err
	}
	return nil
}

func runCompletionZshCmd(cmd *cobra.Command, _ []string) error {
	err := cmd.GenZshCompletion(os.Stdout)
	if err != nil {
		os.Exit(1)
		return err
	}
	return nil
}

func runCompletionFishCmd(cmd *cobra.Command, _ []string) error {
	err := cmd.GenFishCompletion(os.Stdout, true)
	if err != nil {
		os.Exit(1)
		return err
	}
	return nil
}

package cmd

import (
	"os"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCompletionCmd())
}

func newCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates shell completion scripts",
		Long:  "Generates shell completion scripts",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return constants.ErrInvalidParam
		},
	}
	cmd.AddCommand(
		newCompletionBashCmd(),
		newCompletionZshCmd(),
		newCompletionFishCmd(),
	)
	return cmd
}

func newCompletionBashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Long:  "Generates bash completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			err := rootCmd.GenBashCompletion(os.Stdout)
			if err != nil {
				os.Exit(1)
			}
		},
	}
	return cmd
}

func newCompletionZshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Long:  "Generates zsh completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			err := rootCmd.GenZshCompletion(os.Stdout)
			if err != nil {
				os.Exit(1)
			}
		},
	}
	return cmd
}

func newCompletionFishCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fish",
		Short: "Generates fish completion scripts",
		Long:  "Generates fish completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			err := rootCmd.GenFishCompletion(os.Stdout, true)
			if err != nil {
				os.Exit(1)
			}
		},
	}
	return cmd
}

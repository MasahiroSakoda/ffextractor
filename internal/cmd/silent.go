package cmd

import (
	"github.com/spf13/cobra"
)

var silentCmd = &cobra.Command{
	Use:   "silent",
	Short: "Extract media exclude silent parts.",
	Long:  "Extract media exclude silent parts.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

package cmd

import (
	"github.com/spf13/cobra"
)

var blackoutCmd = &cobra.Command{
	Use:   "blackout",
	Short: "Extract video exclude blackout parts.",
	Long:  "Extract video exclude blackout parts.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

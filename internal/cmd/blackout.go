package cmd

import (
	"path/filepath"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newBlackoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "blackout",
		Short: "Extract video exclude blackout parts.",
		Long:  "Extract video exclude blackout parts.",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			path := filepath.Clean(args[0])
			if !util.IsMediaPath(path) || !util.Exists(path) {
				return constants.ErrInvalidParam
			}
			logrus.Debugf("%s", args[0])

			return nil
		},
	}
}

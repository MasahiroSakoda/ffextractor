package cmd

import (
	"os"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var blackoutCmd = &cobra.Command{
	Use:   "blackout",
	Short: "Extract video exclude blackout parts.",
	Long:  "Extract video exclude blackout parts.",
	Args:  cobra.MinimumNArgs(1),
	RunE: extractExcludeBlackout,
}

func extractExcludeBlackout(_ *cobra.Command, args []string) error {
	var path = ""
	contains, err := util.ContainsMedia(args[0])
	if err != nil {
		return constants.ErrInvalidParam
	}

	if contains {
		path = args[0]
	} else {
		os.Exit(1)
		return constants.ErrInvalidParam
	}
	logrus.Debugf("%s", path)

	return nil
}

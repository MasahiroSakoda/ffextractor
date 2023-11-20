package cmd

import (
	"os"
	"fmt"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"
	"github.com/MasahiroSakoda/ffextractor/internal/ui"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
)

var silentCmd = &cobra.Command{
	Use:   "silent",
	Short: "Extract media exclude silent parts.",
	Long:  "Extract media exclude silent parts.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		path := ""
		contains, err := util.ContainsMedia(args[0])
		if err != nil {
			logrus.Errorf("%s: %v", constants.ErrFileNotFound, err)
			return err
		}
		if contains {
			path = args[0]
		} else {
			logrus.Errorf("%s: %v", constants.ErrInvalidParam, err)
			os.Exit(1)
		}
		// segments, err := ffmpeg.DetectSilence(path)
		// if err != nil {
		// 	return err
		// }
		// jsonData, err := json.Marshal(segments)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("%s", string(jsonData))

		m := ui.New(path)
		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			logrus.Errorf("%s: %v", constants.ErrUnexpected, err)
			os.Exit(1)
		}
		p.Quit()
		return nil
	},
}

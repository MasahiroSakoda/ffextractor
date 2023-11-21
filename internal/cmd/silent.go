package cmd

import (
	"os"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/ui"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newSilentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "silent",
		Short: "Extract media exclude silent parts.",
		Long:  "Extract media exclude silent parts.",
		Args:  cobra.MinimumNArgs(1),
		RunE:  func (_ *cobra.Command, args []string) error {
			contains, err := util.ContainsMedia(args[0])
			if err != nil {
				return constants.ErrInvalidParam
			}

			if !contains {
				os.Exit(1)
				return constants.ErrInvalidParam
			}

			m := ui.New(args[0])
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				logrus.Errorf("%s: %v", constants.ErrUnexpected, err)
				os.Exit(1)
			}
			p.Quit()

			return nil
		},
	}
}

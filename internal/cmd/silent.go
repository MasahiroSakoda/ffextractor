package cmd

import (
	"path/filepath"

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
			path := filepath.Clean(args[0])
			if !util.IsMediaPath(path) || !util.Exists(path) {
				return constants.ErrInvalidParam
			}

			m := ui.New(path)
			p := tea.NewProgram(m)
			if _, err := p.Run(); err != nil {
				logrus.Errorf("%s: %v", constants.ErrUnexpected, err)
				return err
			}
			p.Quit()

			return nil
		},
	}
}

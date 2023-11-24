package cmd

import (
	"strconv"
	"strings"

	c "github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/config"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/spf13/cobra"
)

func newConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Configure " + c.CommandName + " options.",
		Long:  "Configure " + c.CommandName + " options.",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(_ *cobra.Command, args []string) error {
			field := strings.ToLower(args[0])
			value := args[1]

			root       := config.Root
			fileCfg    := root.File
			extractCfg := root.Extract
			encodeCfg  := root.Encode

			switch field {
			case c.ConfigOverwrite:
				overwrite, err := util.ParseBoolean([]byte(value))
				if err != nil {
					return err
				}
				fileCfg.Overwrite = overwrite
			case c.ConfigAnnotation:
				if len(value) == 0 {
					return c.ErrInvalidParam
				}
				fileCfg.Annotation = value
			case c.ConfigThreshold:
				threshold, err := util.ParseInteger([]byte(value))
				if err != nil {
					return err
				}
				extractCfg.Threshold = threshold
			case c.ConfigSilenceDuration:
				duration, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				extractCfg.SilenceDuration = duration
			case c.ConfigBlackoutDuration:
				duration, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				extractCfg.BlackoutDuration = duration
			case c.ConfigSplitWithEncode:
				encode, err := util.ParseBoolean([]byte(value))
				if err != nil {
					return err
				}
				encodeCfg.SplitWithEncode = encode
			case c.ConfigConcatWithEncode:
				encode, err := util.ParseBoolean([]byte(value))
				if err != nil {
					return err
				}
				encodeCfg.SplitWithEncode = encode
			}
			configPath, err := util.GetConfigFilePath()
			if err != nil {
				return err
			}
			err = root.Save(configPath)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

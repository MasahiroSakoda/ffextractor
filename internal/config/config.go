package config

import (
	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Load returns result of loading config
func Load() (*Config, error) {
	configDir,  _  := util.GetConfigDir()
	configPath, _  := util.GetConfigFilePath()
	viper.SetConfigName(constants.DefaultConfigFileName)
	viper.SetConfigType(constants.DefaultConfigFileType)
	viper.AddConfigPath(configDir)

	if !util.Exists(configPath) {
		// configure default values
		viper.SetDefault("file", map[string]any{
			"overwrite":  constants.DefaultOverwrite,
			"annotation": constants.DefaultAnnotation,
		})
		viper.SetDefault("extract", map[string]any{
			"threshold":         constants.DefaultThreshold,
			"silence_duration":  constants.DefaultSilenceDuration,
			"blackout_duration": constants.DefaultBlackoutDuration,
		})
		viper.SetDefault("encode", map[string]bool{
			"split_with_encode":  constants.DefaultSplitWithEncode,
			"concat_with_encode": constants.DefaultConcatWithEncode,
		})

		err := viper.SafeWriteConfigAs(configPath)
		if err != nil {
			logrus.Errorf("%s: %v", constants.ErrSaveConfig, err)
			return nil, err
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("%s: %v", constants.ErrLoadConfig, err)
		return nil, err
	}

	var config *Config
	if err := viper.Unmarshal(&config); err != nil {
		logrus.Errorf("%s: %v", constants.ErrParseConfig, err)
		return nil, err
	}
	return config, nil
}

// Save returns result of saving config
func (c *Config) Save() error {
	viper.Set("file",    c.File)
	viper.Set("extract", c.Extract)
	viper.Set("encode",  c.Encode)
	return viper.SafeWriteConfig()
}

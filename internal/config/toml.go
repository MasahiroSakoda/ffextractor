package config

import (
	"io"
	"os"
	"errors"

	"github.com/pelletier/go-toml/v2"

	"github.com/MasahiroSakoda/ffextractor/internal/util"
)

// Root :
var Root = &Config{}

var defaultConfig = &Config{
	File: FileSection{
		Overwrite:  false,
		Annotation: "_merged",
	},
	Extract: ExtractSection{
		Threshold:       -10,
		SilenceDuration:  10.000,
		BlackoutDuration: 10.000,
	},
	Encode: EncodeSection{
		SplitWithEncode:  true,
		ConcatWithEncode: true,
	},
}

func init() {
	config  := defaultConfig
	path, _ := util.GetConfigFilePath()
	exists  := util.Exists(path)
	if !exists {
		if _, err := util.CreateFile(path); err == nil {
			_ = config.Save(path)
		}
	}
}

// Save :
func (config *Config) Save(path string) error {
	buf, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(path, buf, 0644)
}

// Load :
func (config *Config) Load(path string) (*Config, error) {
	var configFile string
	if !util.Exists(path) {
		configFile, _ = util.GetConfigFilePath()
	} else {
		configFile = path
	}
	fh, err := os.Open(configFile)
	if err != nil {
		return config, errors.New("failed to open file")
	}
	defer fh.Close()

	return config, config.Import(fh)
}

// Import :
func (config *Config) Import(fh io.Reader) error {
	data, err := io.ReadAll(fh)
	if err != nil {
		return err
	}
	return toml.Unmarshal(data, &config)
}

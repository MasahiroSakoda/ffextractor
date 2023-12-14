// Package config provides function to use config
package config

import (
	"testing"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/stretchr/testify/assert"
	// "github.com/spf13/viper"
)

func TestLoad(t *testing.T) {
	tests := []struct{
		name    string
		existed bool
	}{
		{ existed: false, name: "Successfully loaded default config values" },
		{ existed: true,  name: "Successfully loaded existed config values" },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			configPath, _ := util.GetConfigFilePath()
			if tt.existed {
				if util.Exists(configPath) {
					err := util.RemoveFile(configPath)
					if err != nil {
						t.Errorf("%s: %v", constants.ErrFileRemove, err)
					}
				}
			}
			c, err := Load()
			if err != nil {
				t.Errorf("%s: %v", constants.ErrLoadConfig, err)
			}
			f, ext, enc := c.File, c.Extract, c.Encode
			assert.Equal(t, constants.DefaultOverwrite,         f.Overwrite)
			assert.Equal(t, constants.DefaultAnnotation,        f.Annotation)
			assert.Equal(t, constants.DefaultThreshold,         ext.Threshold)
			assert.Equal(t, constants.DefaultSilenceDuration,   ext.SilenceDuration)
			assert.Equal(t, constants.DefaultBlackoutDuration,  ext.BlackoutDuration)
			assert.Equal(t, constants.DefaultSplitWithEncode,   enc.SplitWithEncode)
			assert.Equal(t, constants.DefaultConcatWithEncode,  enc.ConcatWithEncode)
		})
	}
}

func TestSave(t *testing.T) {
	tests := []struct{
		name    string
		config  Config
		wantErr bool
	}{
		{
			config: Config{
				File: FileSection{
					Overwrite: true,
					Annotation: "_hoge",
				},
				Extract: ExtractSection{
					Threshold: -50,
					SilenceDuration: 30,
					BlackoutDuration: 30,
				},
				Encode: EncodeSection{
					SplitWithEncode: true,
					ConcatWithEncode: true,
				},
			},
			name: "Successfully saved config with correct value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := Load()
			if err != nil {
				t.Errorf("%s: %v", constants.ErrLoadConfig, err)
			}
			cFile, cExt, cEnc := c.File, c.Extract, c.Encode
			tFile, tExt, tEnc := tt.config.File, tt.config.Extract, tt.config.Encode
			cFile.Overwrite  = tFile.Overwrite
			cFile.Annotation = tFile.Annotation
			cExt.Threshold        = tExt.Threshold
			cExt.SilenceDuration  = tExt.SilenceDuration
			cExt.BlackoutDuration = tExt.BlackoutDuration
			cEnc.SplitWithEncode  = tEnc.SplitWithEncode
			cEnc.ConcatWithEncode = tEnc.ConcatWithEncode
			err = c.Save()
			if err != nil {
				t.Errorf("%s: %v", constants.ErrFileWrite, err)
			}
			assert.Equal(t, cFile.Overwrite, tFile.Overwrite)
			assert.Equal(t, cFile.Overwrite, tFile.Overwrite)
			assert.Equal(t, cExt.Threshold,        tExt.Threshold)
			assert.Equal(t, cExt.SilenceDuration,  tExt.SilenceDuration)
			assert.Equal(t, cExt.BlackoutDuration, tExt.BlackoutDuration)
			assert.Equal(t, cEnc.SplitWithEncode,  tEnc.SplitWithEncode)
			assert.Equal(t, cEnc.ConcatWithEncode, tEnc.ConcatWithEncode)
		})
	}
}

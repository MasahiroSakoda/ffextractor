package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/MasahiroSakoda/ffextractor/internal/util"
)

var (
	c *Config
)

func init() {
	c = defaultConfig
}

func TestSave(t *testing.T) {
	confPath, _ := util.GetConfigFilePath()
	tests := []struct {
		name   string
		path1  string
		path2  string
		isSame bool
	}{
		{ name: "", path1: confPath, path2: "./testdata/config/config.toml", isSame: true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Save(tt.path1)
			if err != nil {
				assert.Error(t, err)
			}
			c1 := c
			err = c.Load(tt.path2)
			if err != nil {
				assert.Error(t, err)
			}
			c2 := c
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.isSame, cmp.Equal(c1, c2))
		})
	}
}

func TestLoad(t *testing.T) {
	confPath, _ := util.GetConfigFilePath()
	tests := []struct {
		name   string
		path1  string
		path2  string
		isSame bool
	}{
		{ name: "", path1: confPath, path2: "./testdata/config/config.toml", isSame: true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := c.Load(tt.path1)
			if err != nil {
				assert.Error(t, err)
			}
			c1 := c
			err = c.Load(tt.path2)
			if err != nil {
				assert.Error(t, err)
			}
			c2 := c
			if err != nil {
				assert.Error(t, err)
			}
			assert.Equal(t, tt.isSame, cmp.Equal(c1, c2))
		})
	}
}

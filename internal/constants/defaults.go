// Package constants provides common constants
package constants

const (
	// CommandName is to use as binary name
	CommandName string = "ffextractor"

	// DefaultConfigFileName is name of config file (without extension)
	DefaultConfigFileName string = "config"
	// DefaultConfigFileType is extension of config file
	DefaultConfigFileType string = "toml"
	// DefaultConfigFileDir is default config file dir
	DefaultConfigFileDir  string = "$HOME/.config/" + CommandName
)

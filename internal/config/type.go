// Package config provides function to use config
package config

// Config : config root
type Config struct {
	File    FileSection    `mapstructure:"file"`
	Extract ExtractSection `mapstructure:"extract"`
	Encode  EncodeSection  `mapstructure:"encode"`
}
// FileSection : file options
type FileSection struct {
	Overwrite  bool   `mapstructure:"overwrite"`
	Annotation string `mapstructure:"annotation"`
}

// ExtractSection : ffmpeg extract options
type ExtractSection struct {
	Threshold        int     `mapstructure:"threshold"`
	SilenceDuration  float64 `mapstructure:"silence_duration"`
	BlackoutDuration float64 `mapstructure:"blackout_duration"`
}

// EncodeSection : ffmpeg encode options
type EncodeSection struct {
	SplitWithEncode  bool `mapstructure:"split_with_encode"`
	ConcatWithEncode bool `mapstructure:"concat_with_encode"`
}

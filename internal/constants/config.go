// Package constants provides common constants
package constants

const (
	// ConfigOverwrite is mapstructure key `overwrite` option
	ConfigOverwrite  = "overwrite"
	// ConfigAnnotation is mapstructure key `annotation` option
	ConfigAnnotation = "annotation"

	// ConfigThreshold is mapstructure key `threshold` option
	ConfigThreshold        = "threshold"
	// ConfigSilenceDuration is mapstructure key `silence_duration` option
	ConfigSilenceDuration  = "silence_duration"
	// ConfigBlackoutDuration is mapstructure key `blackout_duration` option
	ConfigBlackoutDuration = "blackout_duration"

	// ConfigSplitWithEncode is mapstructure key `split_with_encode` option
	ConfigSplitWithEncode  = "split_with_encode"
	// ConfigConcatWithEncode is mapstructure key `concat_with_encode` option
	ConfigConcatWithEncode = "concat_with_encode"
)

const (
	// DefaultOverwrite is default value for `overwrite` option
	DefaultOverwrite  = false
	// DefaultAnnotation is default value for `annotation` option
	DefaultAnnotation = "_merged"

	// DefaultThreshold is default value for `threshold` option
	DefaultThreshold        = -10
	// DefaultSilenceDuration is default value for `silence_duration` option
	DefaultSilenceDuration  = 10.0
	// DefaultBlackoutDuration is default value for `blackout_duration` option
	DefaultBlackoutDuration = 10.0

	// DefaultSplitWithEncode is default value for `split_with_encode` option
	DefaultSplitWithEncode  = true
	// DefaultConcatWithEncode is default value for `concat_with_encode` option
	DefaultConcatWithEncode = true
)

package constants

const (
	ConfigOverwrite  = "overwrite"
	ConfigAnnotation = "annotation"

	ConfigThreshold        = "threshold"
	ConfigSilenceDuration  = "silence_duration"
	ConfigBlackoutDuration = "blackout_duration"

	ConfigSplitWithEncode  = "split_with_encode"
	ConfigConcatWithEncode = "concat_with_encode"
)

const (
	DefaultOverwrite  = false
	DefaultAnnotation = "_merged"

	DefaultThreshold        = -10
	DefaultSilenceDuration  = 10.0
	DefaultBlackoutDuration = 10.0

	DefaultSplitWithEncode  = true
	DefaultConcatWithEncode = true
)

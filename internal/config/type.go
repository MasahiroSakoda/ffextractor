package config

// Config : config root
type Config struct {
	File    FileSection    `toml:"file"`
	Extract ExtractSection `toml:"extract"`
	Encode  EncodeSection  `toml:"encode"`
}
// FileSection : file options
type FileSection struct {
	Overwrite  bool   `toml:"overwrite"`
	Annotation string `toml:"annotation"`
}

// ExtractSection : ffmpeg extract options
type ExtractSection struct {
	Threshold        int64   `toml:"threshold"`
	SilenceDuration  float64 `toml:"silence_duration"`
	BlackoutDuration float64 `toml:"blackout_duration"`
}

// EncodeSection : ffmpeg encode options
type EncodeSection struct {
	SplitWithEncode  bool `toml:"split_with_encode"`
	ConcatWithEncode bool `toml:"concat_with_encode"`
}

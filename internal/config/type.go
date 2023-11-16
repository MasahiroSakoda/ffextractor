package config

// Config : config root
type Config struct {
	File    FileSection    `toml:"file"`
	Extract ExtractSection `toml:"extract"`
}
// File : file options
type FileSection struct {
	Overwrite  bool   `toml:"overwrite"`
	Annotation string `toml:"annotation"`
}

// Extract : ffmpeg extract options
type ExtractSection struct {
	Threshold        int     `toml:"threshold"`
	SilenceDuration  float64 `toml:"silence_duration"`
	BlackoutDuration float64 `toml:"blackout_duration"`
}

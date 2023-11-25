package constants

import "errors"

var (
	// ErrUnexpected is used for unknown reason
	ErrUnexpected = errors.New("unexpected error")

	// ErrSilenceDetect is error for detect silence segments whilte using ffmpeg
	ErrSilenceDetect  = errors.New("silence detection error")
	// ErrBlackoutDetect is error for detect blackout segments whilte using ffmpeg
	ErrBlackoutDetect = errors.New("blackout detection error")
	// ErrParseSegment is error for parse segments
	ErrParseSegment   = errors.New("parse processing error")
	// ErrSplitSegment is error for split segments
	ErrSplitSegment   = errors.New("split processing error")

	// ErrLoadConfig is error for loading config file
	ErrLoadConfig       = errors.New("failed to load config file")
	// ErrSaveConfig is error for loading config file
	ErrSaveConfig       = errors.New("failed to save config file")
	// ErrParseConfig is error for parsing config file
	ErrParseConfig      = errors.New("failed to parse env to config struct")
	// ErrInvalidConfigKey is error for invalid key into config
	ErrInvalidConfigKey = errors.New("invalid config key")

	// ErrFileOpen is error while opening file
	ErrFileOpen     = errors.New("failed to open file")
	// ErrFileRead is error while reading file
	ErrFileRead     = errors.New("failed to read file")
	// ErrFileWrite is error while writing file
	ErrFileWrite    = errors.New("failed to write file")
	// ErrFileRemove is error while removing file
	ErrFileRemove   = errors.New("failed to remove file or directory")
	// ErrFileNotFound is error file existence
	ErrFileNotFound = errors.New("no such file or directory")

	// ErrMkdir is error while making directory
	ErrMkdir      = errors.New("failed to create directory")
	// ErrFileScan is error while scanning file or directory
	ErrFileScan   = errors.New("failed to scan file or directory")

	// ErrJSONEncode is error while encoding JSON
	ErrJSONEncode = errors.New("failed to encode json")
	// ErrJSONDecode is error while decoding JSON
	ErrJSONDecode = errors.New("failed to decode json")
	// ErrTomlEncode is error while encoding TOML
	ErrTomlEncode = errors.New("failed to encode toml")
	// ErrTomlDecode is error while decoding TOML
	ErrTomlDecode = errors.New("failed to decode toml")

	// ErrInvalidParam is error for wrong command line parameters
	ErrInvalidParam = errors.New("invalid parameter")
)

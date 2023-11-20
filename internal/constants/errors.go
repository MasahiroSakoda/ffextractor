package constants

import "errors"

var (
	// ErrUnexpected : unexpected error
	ErrUnexpected = errors.New("unexpected error")

	// Extraction
	ErrSilenceDetect  = errors.New("silence detection error")
	ErrBlackoutDetect = errors.New("blackout detection error")
	ErrSplitSegment   = errors.New("split processing error")
	ErrParseSegment   = errors.New("parse processing error")

	// Config
	ErrLoadConfig       = errors.New("failed to load config file")
	ErrParseConfig      = errors.New("failed to parse env to config struct")
	ErrInvalidConfigKey = errors.New("invalid config key")

	// ErrFileOpen : error for opening file
	ErrFileOpen     = errors.New("failed to open file")
	ErrFileRead     = errors.New("failed to read file")
	ErrFileWrite    = errors.New("failed to write file")
	ErrFileRemove   = errors.New("failed to remove file or directory")
	ErrFileNotFound = errors.New("no such file or directory")

	ErrMkdir      = errors.New("failed to create directory")
	ErrFileScan   = errors.New("failed to scan file or directory")

	ErrJSONEncode = errors.New("failed to encode json")
	ErrJSONDecode = errors.New("failed to decode json")
	ErrTomlEncode = errors.New("failed to encode json")
	ErrTomlDecode = errors.New("failed to decode json")

	ErrInvalidParam = errors.New("invalid parameter")
)

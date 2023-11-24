package util

import (
	"io/fs"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	rootDir = filepath.Join(filepath.Dir(b), "../..")
)

func TestExists(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     bool
	}{
		{ name: "Find file exists", filepath: "file.go", want: true },
		{ name: "Return false for missing file", filepath: "afilethatdoesntexist.go", want: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Exists(tt.filepath)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAudioPath(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		expect bool
	}{
		{ path: "example.mp3",  expect: true,  name: "Successful return with mp3 file" },
		{ path: "example.m4a",  expect: true,  name: "Successful return with m4a file" },
		{ path: "example.ogg",  expect: true,  name: "Successful return with ogg file" },
		{ path: "example.oga",  expect: true,  name: "Successful return with oga file" },
		{ path: "example.wav",  expect: true,  name: "Successful return with wav file" },
		{ path: "example.aif",  expect: true,  name: "Successful return with aif file" },
		{ path: "example.aiff", expect: true,  name: "Successful return with aiff file" },
		{ path: "example.flac", expect: true,  name: "Successful return with flac file" },
		{ path: "example.m3p",  expect: false, name: "Failed with wrong file" },
		{ path: "example.fla",  expect: false, name: "Failed with wrong file" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAudioPath(tt.path)
			assert.Equal(t, tt.expect, got)
		})
	}
}
func TestVideoPath(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		expect bool
	}{
		{ path: "example.mp4",  expect: true,  name: "Successful return with mp4 file" },
		{ path: "example.mkv",  expect: true,  name: "Successful return with mkv file" },
		{ path: "example.mov",  expect: true,  name: "Successful return with mov file" },
		{ path: "example.webm", expect: true,  name: "Successful return with webm file" },
		{ path: "example.mpg",  expect: true,  name: "Successful return with mpg file" },
		{ path: "example.avi",  expect: true,  name: "Successful return with avi file" },
		{ path: "example.flv",  expect: true,  name: "Successful return with flv file" },
		{ path: "example.wmv",  expect: true,  name: "Successful return with wmv file" },
		{ path: "example.m4p",  expect: false, name: "Failed with wrong file" },
		{ path: "example.web",  expect: false, name: "Failed with wrong file" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsVideoPath(tt.path)
			assert.Equal(t, tt.expect, got)
		})
	}
}

func TestContainsMedia(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		expect bool
	}{
		{ path: rootDir + "/testdata/ffmpeg", expect: true,  name: "Successful return with proper path" },
		{ path: rootDir + "/testdata/config", expect: false, name: "Failed with path does not contain media file" },
		{ path: rootDir + "/completions",     expect: false, name: "Failed with path does not contain media file" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsMedia(tt.path)
			assert.Equal(t, tt.expect, got)
		})
	}
}
func TestIsExecutable(t *testing.T) {
	tests := []struct {
		perm fs.FileMode
		want bool
	}{
		{0101, false},
		{0111, true},
		{0644, false},
		{0666, false},
		{0777, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.perm), func(t *testing.T) {
			tmp, err := os.CreateTemp(os.TempDir(), "slides-*")
			if err != nil {
				t.Fatal("failed to create temp file")
			}
			defer os.Remove(tmp.Name())
			err = tmp.Chmod(tt.perm)
			if err != nil {
				t.Fatal(err)
			}
			s, err := tmp.Stat()
			if err != nil {
				t.Fatal("failed to stat")
			}

			want := tt.want
			got  := IsExecutable(s)
			if tt.want != got {
				t.Log(want)
				t.Log(got)
				t.Fatalf("IsExecutable returned an incorrect result, want: %t, got %t", want, got)

			}
		})
	}
}

func TestUnixHomeDir(t *testing.T) {
	tests := []struct {
		name         string
		wantErr      bool
	}{
		{ name: "Successful fetch of $HOME directory", wantErr: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := UnixHomeDir()
			if (err != nil) != tt.wantErr {
				t.Fatalf("UnixHomeDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetConfigDir(t *testing.T) {
	tests := []struct {
		name         string
		wantErr      bool
	}{
		{ name: "Successful fetch of $XDG_CONFIG_HOME directory", wantErr: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetConfigDir()
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetConfigDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFilenameFromPath(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		file   string
		expect bool
	}{
		{ path: "./valid_file.txt",       file: "valid_file.txt", expect: true,  name: "Successful return filename with relative path" },
		{ path: "./invalid_file.txt",     file: "valid_file.txt", expect: false, name: "Fail with wrong relative path" },
		{ path: "~/dev/valid_file.txt",   file: "valid_file.txt", expect: true,  name: "Successful return filename with relative path" },
		{ path: "~/dev/invalid_file.txt", file: "valid_file.txt", expect: false, name: "Fail with wrong relative path" },
		{ path: "/tmp/valid_file.txt",    file: "valid_file.txt", expect: true,  name: "Successful return filename with absolute path" },
		{ path: "/tmp/invalid_file.txt",  file: "valid_file.txt", expect: false, name: "Fail with wrong absolute path" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFilenameFromPath(tt.path)
			assert.Equal(t, tt.expect, got == tt.file )
		})
	}
}

func TestGetConfigFilePath(t *testing.T) {
	tests := []struct {
		name         string
		wantErr      bool
	}{
		{ name: "Successful fetch of config.toml", wantErr: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetConfigFilePath()
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetConfigFilePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetFileList(t *testing.T) {
	homeDir, _ := UnixHomeDir()
	tests := []struct {
		name     string
		path     string
		wantErr  bool
	}{
		{ name: "Successful", path: homeDir, wantErr: false },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileList(tt.path)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GetFileList() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, tt.wantErr, got == nil)
		})
	}
}

// TODO: implement test
func TestGetFileListByExts(t *testing.T) {
}

// TODO: implement test
func TestGetFileListByRegexp(t *testing.T) {
}

func TestCreateFile(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{ name: "Create proper file", filepath: "create.txt", wantErr: false },
		{ name: "Return false without filename", filepath: "", wantErr: true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFile(tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, want %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantErr, got == nil)
			os.Remove(tt.filepath)
		})
	}
}

func TestRemoveFile(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{ name: "Remove existed file", filepath: "remove.txt", wantErr: false },
		{ name: "Remove missing file", filepath: "", wantErr: true },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, errCreate := CreateFile(tt.filepath); errCreate != nil {
				assert.Equal(t, tt.wantErr, errCreate != nil)
			}
			if errRemove := RemoveFile(tt.filepath); errRemove != nil {
				assert.Equal(t, tt.wantErr, errRemove != nil)
			}
		})
	}
}

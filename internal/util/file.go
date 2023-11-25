package util

import (
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"errors"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
)

// Exists returns file existence
func Exists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsAudioPath :
func IsAudioPath(path string) bool {
	for _, ext := range constants.AudioExtensions {
		if strings.HasSuffix(path, "." + ext) {
			return true
		}
	}
	return false
}

// IsVideoPath :
func IsVideoPath(path string) bool {
	for _, ext := range constants.VideoExtensions {
		if strings.HasSuffix(path, "." + ext) {
			return true
		}
	}
	return false
}

// IsMediaPath :
func IsMediaPath(path string) bool {
	return IsAudioPath(path) || IsVideoPath(path)
}

// ContainsMedia :
func ContainsMedia(path string) bool {
	isMedia := false
	_ = filepath.WalkDir(path, func(p  string, d fs.DirEntry, err error) error {
		info, _ := d.Info()
		if !info.IsDir() && IsMediaPath(p) {
			isMedia = true
		}
		return nil
	})
	return isMedia
}

// IsExecutable returns whether a file has execution permissions
func IsExecutable(s fs.FileInfo) bool {
	return s.Mode().Perm()&0111 == 0111
}

// UnixHomeDir returns $HOME directory
func UnixHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return usr.HomeDir, err
	}
	return os.Getenv("HOME"), nil
}

// GetConfigDir returns $XDG_CONFIG_HOME directory
func GetConfigDir() (string, error) {
	if userConfigDir, err := os.UserConfigDir(); err != nil {
		return filepath.Join(userConfigDir, constants.CommandName), nil
	}
	homeDir, _ := UnixHomeDir()
	if homeDir == "" {
		return "", errors.New("unable to get current user home directory: os/user lookup failed; $HOME is empty")
	}
	configDir := filepath.Join(homeDir, ".config", constants.CommandName)
	if !Exists(configDir) {
		if err := os.Mkdir(configDir, 0600); err != nil {
			return "", constants.ErrFileRead
		}
	}
	return configDir, nil
}

// GetConfigFilePath returns config file path for `ffextractor`
func GetConfigFilePath() (string, error) {
	configFile := "config.toml"
	configDir, err := GetConfigDir()
	if err != nil {
		return filepath.Join("~/.config/ffextractor/", configFile), nil
	}
	return filepath.Join(configDir, configFile), nil
}

// GetFilenameFromPath returns filename from path
func GetFilenameFromPath(path string) string {
	segments := strings.Split(path, "/")
	return segments[len(segments) - 1]
}

// GetFileList returns file list
func GetFileList(p string) ([]string, error) {
	return filepath.Glob(p)
}

// GetFileListByExts returns file list filter by extension
func GetFileListByExts(dir string, exts []string, threads int) ([]string, error) {
	files := make([]string, 512)
	ch    := make(chan string, threads)
	done  := make(chan int)
	go func() {
		for file := range ch {
			files = append(files, file)
		}
		done <- 1
	}()

	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		for _, ext := range exts {
			if !info.IsDir() && strings.HasSuffix(path, ext){
				ch <- filepath.Join(dir, path)
			}
		}
		return err
	})
	close(ch)
	<-done
	if err != nil {
		return nil, err
	}
	return files, nil
}

// GetFileListByRegexp returns file list filter by regexp
func GetFileListByRegexp(dir string, pattern *regexp.Regexp, threads int) ([]string, error) {
	files := make([]string, 512)
	ch    := make(chan string, threads)
	done  := make(chan int)
	go func() {
		for file := range ch {
			files = append(files, file)
		}
		done <- 1
	}()

	err := filepath.WalkDir(dir, func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() && pattern.MatchString(info.Name()) {
			ch <- filepath.Join(dir, path)
		}
		return err
	})
	close(ch)
	<-done
	if err != nil {
		return nil, err
	}
	return files, nil
}

// CreateFile create file with file path
func CreateFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), os.ModePerm); err != nil {
		return nil, err
	}
	return os.Create(p)
}

// RemoveFile remove file with file path
func RemoveFile(p string) error {
	var err = os.Remove(p)

	if err != nil {
		return err
	}
	return nil
}

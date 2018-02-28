package fileutils

import (
	"errors"
	"os"
	"path/filepath"
)

var (
	ErrNotDirectory = errors.New("PathError must give path to directory not file")
)

// DirSize takes a directory path, then returns the total size of all files and error
func DirSize(path string) (size int64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

func DirFilenames(path string) (info []string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	// If Not Directory return error
	if !fi.IsDir() {
		return nil, ErrNotDirectory
	}
	defer f.Close()
	return f.Readdirnames(0)
}

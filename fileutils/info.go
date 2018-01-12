package fileutils

import (
	"os"
	"path/filepath"
)

// DirectorySize takes a directory path, then returns the total size of all files and error
func DirectorySize(path string) (size int64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

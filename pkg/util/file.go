package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilePath(path string) ([]string, error) {
	if !strings.HasSuffix(path, ".js") {
		return GetFilePathsInDir(path)
	}
	return []string{path}, nil
}

func GetFilePathsInDir(dir string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".js") {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths, err
}

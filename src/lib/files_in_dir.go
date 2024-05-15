package lib

import (
	"os"
	"path/filepath"
)

func FilesInDir(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, FilesInDir(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

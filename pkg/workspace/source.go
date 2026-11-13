package workspace

import (
	"io/fs"
	"path/filepath"
)

func FindSources() ([]string, error) {
	var paths []string

	err := filepath.Walk("src", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		paths = append(paths, path)
		return nil
	})

	return paths, err
}

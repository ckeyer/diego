package metadata

import (
	"os"
	"path/filepath"
)

func LoadFromDir(dir string) error {
	filepath.Walk(dir, walkFn)
	return nil
}

func walkFn(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	return nil
}

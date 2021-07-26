package dirwalker

import (
	"os"
	"path/filepath"
)

func Walk(root string, fileHashingChannel chan string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileHashingChannel <- path
		}
		return nil
	})
}

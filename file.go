package base

import (
	"os"
	"path/filepath"
	"strings"
)

func FilesHashIndexing(root string) (hash map[string]string) {
	hash = map[string]string{}
	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			n := strings.Replace(path, "\\", "/", -1)
			hash[n] = FileSha512(path)
		}
		return nil
	})
	return
}

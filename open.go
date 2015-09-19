package base

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func OpenUrl(url string) (b []byte) {
	r, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer r.Body.Close()
	b, _ = ioutil.ReadAll(r.Body)

	return
}

func OpenLocal(path string) (b []byte) {
	b, _ = ioutil.ReadFile(path)
	return
}

func SaveLocal(path string, b []byte) {
	dir, _ := filepath.Split(path)
	os.MkdirAll(dir, 0777)
	ioutil.WriteFile(path, b, 0777)
}

func IsExistLocal(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

package base

import (
	"io/ioutil"
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

func FileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func FileTmep(file string, f func() []byte) []byte {
	if FileExist(file) {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		return body
	} else {
		b := f()
		dir, _ := filepath.Split(file)
		os.MkdirAll(dir, 0)
		err := ioutil.WriteFile(file, b, 0666)
		if err != nil {
			panic(err)
		}
		return b
	}
}

func FileOpen(file string) []byte {
	f, _ := ioutil.ReadFile(file)
	return f
}

func FileAppend(file string, w []byte) {
	dir, _ := filepath.Split(file)
	os.MkdirAll(dir, 0)
	ioutil.WriteFile(file, w, os.ModeAppend)
}

func FileWrite(file string, w []byte) {
	dir, _ := filepath.Split(file)
	os.MkdirAll(dir, 0)
	ioutil.WriteFile(file, w, 0666)
}

func FileDel(file string) {
	os.RemoveAll(file)
}

func FileCopy(file string, filenew string) {
	FileWrite(filenew, FileOpen(file))

}

func FileMove(file string, filenew string) {
	FileCopy(file, filenew)
	FileDel(file)
}

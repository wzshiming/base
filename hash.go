package base

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
)

func BytesMd5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
func FileMd5(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return BytesMd5(b)
}
func BytesSha1(b []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(b))
}
func FileSha1(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return BytesSha1(b)
}

func BytesSha256(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

func FileSha256(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return BytesSha256(b)
}

func BytesSha512(b []byte) string {

	return fmt.Sprintf("%x", sha512.Sum512(b))
}

func FileSha512(filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return BytesSha512(b)
}

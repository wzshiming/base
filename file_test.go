package base

import (
	"testing"
)

func Test_file(t *testing.T) {
	INFO(string(EnJson(FilesHashIndexing("./")).Bytes()))
}

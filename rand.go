package base

import (
	"fmt"
	"math/rand"
	"time"
)

var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 147852
var c uint32 = 963852741
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandUint32() uint32 {
	return r.Uint32()
}

func RandString() string {
	return BytesMd5([]byte(fmt.Sprint(RandUint32())))
}

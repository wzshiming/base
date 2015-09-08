package base

import (
	"math/rand"
	"time"
)

var x0 uint32 = uint32(time.Now().UnixNano())
var a uint32 = 147852
var c uint32 = 963852741

var LCG chan uint32

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	LCG = make(chan uint32, 1024)
	go func() {
		for {
			LCG <- r.Uint32()
		}
	}()
}

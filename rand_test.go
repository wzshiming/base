package base

import (
	"testing"
)

func Test_Rand(t *testing.T) {
	for i := 0; i < 100; i++ {
		select {
		case <-LCG:

		default:
			t.Fail()
		}
	}
}

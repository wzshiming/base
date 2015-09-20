package base

import (
	"testing"
)

func Test_stora(t *testing.T) {
	s := NewStora()

	for j := 0; j != 10; j++ {
		i := 0
		s.Get("size", &i)
		i++
		s.Set("size", i)

		t.Log(s.Data())
	}

}

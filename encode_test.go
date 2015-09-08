package base

import (
	"testing"
)

func Test_encodejson(t *testing.T) {

	s := []byte(`{"asdasd":"aaa"}`)
	es := NewEncodeBytes(s)
	m := make(map[string]string)
	es.DeJson(&m)
	es.ReSet()
	es.EnJson(m)
	if string(s) != string(es.Bytes()) {
		t.Fatal(s, es.Bytes())
	}

}

func Test_encodeGob(t *testing.T) {
	es := NewEncodeBytes(nil)
	m := map[string]interface{}{
		"aa": "aaa",
	}
	m2 := make(map[string]interface{})
	es.EnGob(m)
	es.DeGob(&m2)
	t.Log(m2)
}

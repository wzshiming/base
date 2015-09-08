package base

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type EncodeBytes []byte

func NewEncodeBytes(b []byte) *EncodeBytes {
	r := &EncodeBytes{}
	if b != nil {
		r.Set(b)
	}
	return r
}

func (en *EncodeBytes) ReSet() {
	*en = []byte{}
}

func (en *EncodeBytes) Set(b []byte) {
	*en = b
	return
}

func (en *EncodeBytes) Bytes() []byte {
	return []byte(*en)
}

func (en *EncodeBytes) EnJson(s interface{}) {
	*en, _ = json.Marshal(s)
	return
}

func (en *EncodeBytes) DeJson(s interface{}) {
	json.Unmarshal(*en, s)
	return
}

func (en *EncodeBytes) EnGob(s interface{}) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(s); err == nil {
		*en = buf.Bytes()
	}
	return
}

func (en *EncodeBytes) DeGob(s interface{}) {
	buf := bytes.NewBuffer(*en)
	dec := gob.NewDecoder(buf)
	dec.Decode(s)
	return
}

func EnJson(s interface{}) *EncodeBytes {
	b := NewEncodeBytes(nil)
	b.EnJson(s)
	return b
}

func EnGob(s interface{}) *EncodeBytes {
	b := NewEncodeBytes(nil)
	b.EnGob(s)
	return b
}

func SumJson(d ...*EncodeBytes) *EncodeBytes {
	sum := make(map[string]interface{})
	for _, v := range d {
		v.DeJson(&sum)
	}
	return EnJson(sum)
}

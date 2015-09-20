package base

type Stora map[string]*EncodeBytes

func NewStora() *Stora {
	return &Stora{}
}

func (st *Stora) Set(k string, v interface{}) {
	if (*st)[k] == nil {
		(*st)[k] = EnJson(v)
	} else {
		(*st)[k].EnJson(v)
	}
}

func (st *Stora) Get(k string, v interface{}) {
	if (*st)[k] != nil {
		(*st)[k].DeJson(v)
	}
}

func (st *Stora) Del(k string) {
	delete((*st), k)
}

func (st *Stora) Data() (r map[string]interface{}) {
	r = map[string]interface{}{}
	for k, v := range *st {
		var i interface{}
		v.DeJson(&i)
		r[k] = i
	}
	return
}

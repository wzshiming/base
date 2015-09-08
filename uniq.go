package base

import (
	"unsafe"
)

type Unique struct {
	U uint `json:"uniq,String"`
}

func (u *Unique) InitUint() {
	u.U = (uint)((uintptr)(unsafe.Pointer(u)))
}

func (u *Unique) ToUint() uint {
	return u.U
}

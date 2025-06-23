package ssa

import "bytes"

type Bytes struct {
	Bytes []byte
	Liveness
	HasToken
}

func (v *Bytes) Dependencies() []Value {
	return nil
}

func (a *Bytes) Equals(v Value) bool {
	b, sameType := v.(*Bytes)

	if !sameType {
		return false
	}

	return bytes.Equal(a.Bytes, b.Bytes)
}

func (v *Bytes) IsConst() bool {
	return true
}

func (v *Bytes) String() string {
	return string(v.Bytes)
}
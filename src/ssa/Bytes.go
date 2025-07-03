package ssa

import (
	"bytes"
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

type Bytes struct {
	Id
	Bytes []byte
	Liveness
	Source
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

func (v *Bytes) Debug(expand bool) string {
	return v.String()
}

func (v *Bytes) String() string {
	return strconv.Quote(string(v.Bytes))
}

func (v *Bytes) Type() types.Type {
	return types.CString
}
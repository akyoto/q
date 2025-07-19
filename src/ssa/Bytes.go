package ssa

import (
	"bytes"
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

type Bytes struct {
	Bytes     []byte
	Structure *Struct
	Liveness
	Source
}

func (v *Bytes) Inputs() []Value  { return nil }
func (v *Bytes) IsConst() bool    { return true }
func (v *Bytes) String() string   { return strconv.Quote(string(v.Bytes)) }
func (v *Bytes) Struct() *Struct  { return v.Structure }
func (v *Bytes) Type() types.Type { return types.CString }

func (a *Bytes) Equals(v Value) bool {
	b, sameType := v.(*Bytes)

	if !sameType {
		return false
	}

	return bytes.Equal(a.Bytes, b.Bytes)
}
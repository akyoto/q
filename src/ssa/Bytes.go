package ssa

import (
	"bytes"
	"strconv"

	"git.urbach.dev/cli/q/src/types"
)

// Bytes is a raw slice of bytes.
type Bytes struct {
	Bytes []byte
	Independent
	Liveness
	Source
}

// Equals returns true if the byte slices are equal.
func (a *Bytes) Equals(v Value) bool {
	b, sameType := v.(*Bytes)

	if !sameType {
		return false
	}

	return bytes.Equal(a.Bytes, b.Bytes)
}

// IsPure returns true because a byte slice is always constant.
func (b *Bytes) IsPure() bool { return true }

// String returns a human-readable representation of the byte slice.
func (b *Bytes) String() string { return strconv.Quote(string(b.Bytes)) }

// Type returns the CString type.
func (b *Bytes) Type() types.Type { return types.CString }
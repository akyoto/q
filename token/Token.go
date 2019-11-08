package token

import "github.com/akyoto/stringutils/unsafe"

// Token represents a single element in a source file.
// The characters that make up an identifier are grouped into a single token.
// This makes parsing easier and allows us to do better syntax checks.
type Token struct {
	Kind     Kind
	Bytes    []byte
	Position int
}

// String converts the bytes of the token to a string.
func (t Token) String() string {
	return unsafe.BytesToString(t.Bytes)
}

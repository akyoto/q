package token

import (
	"unsafe"
)

// Position is the data type for storing file offsets.
type Position = uint32

// Length is the data type for storing token lengths.
type Length = uint16

// Token represents a single element in a source file.
// The characters that make up an identifier are grouped into a single token.
// This makes parsing easier and allows us to do better syntax checks.
type Token struct {
	Position Position
	Length   Length
	Kind     Kind
}

// Bytes returns the byte slice.
func (t Token) Bytes(buffer []byte) []byte {
	return buffer[t.Position : t.Position+Position(t.Length)]
}

// Start returns the start position.
func (t Token) Start() Position {
	return t.Position
}

// End returns the position after the token.
func (t Token) End() Position {
	return t.Position + Position(t.Length)
}

// Reset resets the token to default values.
func (t *Token) Reset() {
	t.Position = 0
	t.Length = 0
	t.Kind = Invalid
}

// StringFrom returns the token string.
func (t Token) StringFrom(buffer []byte) string {
	return unsafe.String(unsafe.SliceData(t.Bytes(buffer)), t.Length)
}
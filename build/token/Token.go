package token

import "github.com/akyoto/stringutils/unsafe"

// Token represents a single element in a source file.
// The characters that make up an identifier are grouped into a single token.
// This makes parsing easier and allows us to do better syntax checks.
type Token struct {
	Kind     Kind
	Position uint16
	Bytes    []byte
}

// Text converts the bytes of the token to a string.
func (t Token) Text() string {
	return unsafe.BytesToString(t.Bytes)
}

// String might include the kind of token and the token text.
// It is meant to be used for debugging via fmt.Print().
func (t Token) String() string {
	return t.Text()
}

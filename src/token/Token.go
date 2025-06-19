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

// End returns the position after the token.
func (t Token) End() Position {
	return t.Position + Position(t.Length)
}

// IsAssignment returns true if the token is an assignment operator.
func (t Token) IsAssignment() bool {
	return t.Kind > ___ASSIGNMENTS___ && t.Kind < ___END_ASSIGNMENTS___
}

// IsComparison returns true if the token is a comparison operator.
func (t Token) IsComparison() bool {
	return t.Kind > ___COMPARISONS___ && t.Kind < ___END_COMPARISONS___
}

// IsExpressionStart returns true if the token starts an expression.
func (t Token) IsExpressionStart() bool {
	return t.Kind == GroupStart || t.Kind == ArrayStart || t.Kind == BlockStart
}

// IsKeyword returns true if the token is a keyword.
func (t Token) IsKeyword() bool {
	return t.Kind > ___KEYWORDS___ && t.Kind < ___END_KEYWORDS___
}

// IsNumeric returns true if the token is a number or rune.
func (t Token) IsNumeric() bool {
	return t.Kind == Number || t.Kind == Rune
}

// IsOperator returns true if the token is an operator.
func (t Token) IsOperator() bool {
	return t.Kind > ___OPERATORS___ && t.Kind < ___END_OPERATORS___
}

// IsUnaryOperator returns true if the token is a unary operator.
func (t Token) IsUnaryOperator() bool {
	return t.Kind > ___UNARY___ && t.Kind < ___END_UNARY___
}

// Reset resets the token to default values.
func (t *Token) Reset() {
	t.Position = 0
	t.Length = 0
	t.Kind = Invalid
}

// String returns the token string.
func (t Token) String(buffer []byte) string {
	return unsafe.String(unsafe.SliceData(t.Bytes(buffer)), t.Length)
}
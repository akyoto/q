package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	tokens := token.Tokenize([]byte("main(){}"))

	expected := []token.Kind{
		token.Identifier,
		token.GroupStart,
		token.GroupEnd,
		token.BlockStart,
		token.BlockEnd,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestKeyword(t *testing.T) {
	tokens := token.Tokenize([]byte("assert const else extern if import for loop return switch"))

	expected := []token.Kind{
		token.Assert,
		token.Const,
		token.Else,
		token.Extern,
		token.If,
		token.Import,
		token.For,
		token.Loop,
		token.Return,
		token.Switch,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestArray(t *testing.T) {
	tokens := token.Tokenize([]byte("array[i]"))

	expected := []token.Kind{
		token.Identifier,
		token.ArrayStart,
		token.Identifier,
		token.ArrayEnd,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNewline(t *testing.T) {
	tokens := token.Tokenize([]byte("\n\n"))

	expected := []token.Kind{
		token.NewLine,
		token.NewLine,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNumber(t *testing.T) {
	tokens := token.Tokenize([]byte(`123 456`))

	expected := []token.Kind{
		token.Number,
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestOperator(t *testing.T) {
	tokens := token.Tokenize([]byte(`a + b - c * d / e % f << g >> h & i | j ^ k`))

	expected := []token.Kind{
		token.Identifier,
		token.Add,
		token.Identifier,
		token.Sub,
		token.Identifier,
		token.Mul,
		token.Identifier,
		token.Div,
		token.Identifier,
		token.Mod,
		token.Identifier,
		token.Shl,
		token.Identifier,
		token.Shr,
		token.Identifier,
		token.And,
		token.Identifier,
		token.Or,
		token.Identifier,
		token.Xor,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestOperatorAssign(t *testing.T) {
	tokens := token.Tokenize([]byte(`a = b += c -= d *= e /= f %= g &= h |= i ^= j <<= k >>= l`))

	expected := []token.Kind{
		token.Identifier,
		token.Assign,
		token.Identifier,
		token.AddAssign,
		token.Identifier,
		token.SubAssign,
		token.Identifier,
		token.MulAssign,
		token.Identifier,
		token.DivAssign,
		token.Identifier,
		token.ModAssign,
		token.Identifier,
		token.AndAssign,
		token.Identifier,
		token.OrAssign,
		token.Identifier,
		token.XorAssign,
		token.Identifier,
		token.ShlAssign,
		token.Identifier,
		token.ShrAssign,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestOperatorEquality(t *testing.T) {
	tokens := token.Tokenize([]byte(`a == b != c <= d >= e < f > g`))

	expected := []token.Kind{
		token.Identifier,
		token.Equal,
		token.Identifier,
		token.NotEqual,
		token.Identifier,
		token.LessEqual,
		token.Identifier,
		token.GreaterEqual,
		token.Identifier,
		token.Less,
		token.Identifier,
		token.Greater,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestOperatorLogical(t *testing.T) {
	tokens := token.Tokenize([]byte(`a && b || c`))

	expected := []token.Kind{
		token.Identifier,
		token.LogicalAnd,
		token.Identifier,
		token.LogicalOr,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestDefine(t *testing.T) {
	tokens := token.Tokenize([]byte(`a := b`))

	expected := []token.Kind{
		token.Identifier,
		token.Define,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestDot(t *testing.T) {
	tokens := token.Tokenize([]byte(`a.b.c`))

	expected := []token.Kind{
		token.Identifier,
		token.Dot,
		token.Identifier,
		token.Dot,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNot(t *testing.T) {
	tokens := token.Tokenize([]byte(`!a`))

	expected := []token.Kind{
		token.Not,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNegateFirstToken(t *testing.T) {
	tokens := token.Tokenize([]byte(`-a`))

	expected := []token.Kind{
		token.Negate,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNegateAfterGroupStart(t *testing.T) {
	tokens := token.Tokenize([]byte(`(-a)`))

	expected := []token.Kind{
		token.GroupStart,
		token.Negate,
		token.Identifier,
		token.GroupEnd,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNegateSub(t *testing.T) {
	tokens := token.Tokenize([]byte(`-a-b`))

	expected := []token.Kind{
		token.Negate,
		token.Identifier,
		token.Sub,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNegateAfterOperator(t *testing.T) {
	tokens := token.Tokenize([]byte(`-a + -b`))

	expected := []token.Kind{
		token.Negate,
		token.Identifier,
		token.Add,
		token.Negate,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestNegateNumber(t *testing.T) {
	tokens := token.Tokenize([]byte(`-1`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestBinaryNumber(t *testing.T) {
	tokens := token.Tokenize([]byte(`0b1010`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestOctalNumber(t *testing.T) {
	tokens := token.Tokenize([]byte(`0o755`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestHexadecimalNumber(t *testing.T) {
	tokens := token.Tokenize([]byte(`0xCAFE`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestStandaloneZero(t *testing.T) {
	tokens := token.Tokenize([]byte(`0`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestLeadingZero(t *testing.T) {
	tokens := token.Tokenize([]byte(`0123`))

	expected := []token.Kind{
		token.Number,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestRange(t *testing.T) {
	tokens := token.Tokenize([]byte("a..b"))

	expected := []token.Kind{
		token.Identifier,
		token.Range,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestSeparator(t *testing.T) {
	tokens := token.Tokenize([]byte("a,b,c"))

	expected := []token.Kind{
		token.Identifier,
		token.Separator,
		token.Identifier,
		token.Separator,
		token.Identifier,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestComment(t *testing.T) {
	tokens := token.Tokenize([]byte("// Hello\n// World"))

	expected := []token.Kind{
		token.Comment,
		token.NewLine,
		token.Comment,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}

	tokens = token.Tokenize([]byte("// Hello\n"))

	expected = []token.Kind{
		token.Comment,
		token.NewLine,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}

	tokens = token.Tokenize([]byte(`// Hello`))

	expected = []token.Kind{
		token.Comment,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}

	tokens = token.Tokenize([]byte(`//`))

	expected = []token.Kind{
		token.Comment,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}

	tokens = token.Tokenize([]byte(`/`))

	expected = []token.Kind{
		token.Div,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestInvalid(t *testing.T) {
	tokens := token.Tokenize([]byte(`@@`))

	expected := []token.Kind{
		token.Invalid,
		token.Invalid,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestInvalidScript(t *testing.T) {
	tokens := token.Tokenize([]byte(`##`))

	expected := []token.Kind{
		token.Invalid,
		token.Invalid,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestString(t *testing.T) {
	tokens := token.Tokenize([]byte(`"Hello" "World"`))

	expected := []token.Kind{
		token.String,
		token.String,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestStringMultiline(t *testing.T) {
	tokens := token.Tokenize([]byte("\"Hello\nWorld\""))

	expected := []token.Kind{
		token.String,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestStringEOF(t *testing.T) {
	tokens := token.Tokenize([]byte(`"EOF`))

	expected := []token.Kind{
		token.String,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestReturnType(t *testing.T) {
	tokens := token.Tokenize([]byte("()->"))

	expected := []token.Kind{
		token.GroupStart,
		token.GroupEnd,
		token.ReturnType,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestMinusAtEOF(t *testing.T) {
	tokens := token.Tokenize([]byte("1-"))

	expected := []token.Kind{
		token.Number,
		token.Sub,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestRune(t *testing.T) {
	tokens := token.Tokenize([]byte("'a'"))

	expected := []token.Kind{
		token.Rune,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}

func TestScript(t *testing.T) {
	tokens := token.Tokenize([]byte("#!/usr/bin/env q"))

	expected := []token.Kind{
		token.Script,
		token.EOF,
	}

	for i, kind := range expected {
		assert.Equal(t, tokens[i].Kind, kind)
	}
}
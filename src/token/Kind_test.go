package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestTokenGroups(t *testing.T) {
	assert.True(t, token.Assign.IsAssignment())
	assert.True(t, token.Add.IsOperator())
	assert.True(t, token.If.IsKeyword())
	assert.True(t, token.Not.IsUnaryOperator())
	assert.True(t, token.Number.IsNumeric())
	assert.True(t, token.Equal.IsComparison())
}

func TestTokenKindString(t *testing.T) {
	assert.Equal(t, token.Struct.String(), "$")
	assert.Equal(t, token.Dot.String(), ".")
	assert.Equal(t, token.Call.String(), "λ")
	assert.Equal(t, token.Array.String(), "@")
	assert.Equal(t, token.Negate.String(), "-")
	assert.Equal(t, token.Not.String(), "!")
	assert.Equal(t, token.Mul.String(), "*")
	assert.Equal(t, token.Div.String(), "/")
	assert.Equal(t, token.Mod.String(), "%")
	assert.Equal(t, token.Add.String(), "+")
	assert.Equal(t, token.Sub.String(), "-")
	assert.Equal(t, token.Shr.String(), ">>")
	assert.Equal(t, token.Shl.String(), "<<")
	assert.Equal(t, token.And.String(), "&")
	assert.Equal(t, token.Xor.String(), "^")
	assert.Equal(t, token.Or.String(), "|")
	assert.Equal(t, token.Greater.String(), ">")
	assert.Equal(t, token.Less.String(), "<")
	assert.Equal(t, token.GreaterEqual.String(), ">=")
	assert.Equal(t, token.LessEqual.String(), "<=")
	assert.Equal(t, token.Equal.String(), "==")
	assert.Equal(t, token.NotEqual.String(), "!=")
	assert.Equal(t, token.LogicalAnd.String(), "&&")
	assert.Equal(t, token.LogicalOr.String(), "||")
	assert.Equal(t, token.Range.String(), "..")
	assert.Equal(t, token.Separator.String(), ",")
	assert.Equal(t, token.Assign.String(), "=")
	assert.Equal(t, token.Define.String(), ":=")
	assert.Equal(t, token.AddAssign.String(), "+=")
	assert.Equal(t, token.SubAssign.String(), "-=")
	assert.Equal(t, token.MulAssign.String(), "*=")
	assert.Equal(t, token.DivAssign.String(), "/=")
	assert.Equal(t, token.ModAssign.String(), "%=")
	assert.Equal(t, token.AndAssign.String(), "&=")
	assert.Equal(t, token.OrAssign.String(), "|=")
	assert.Equal(t, token.XorAssign.String(), "^=")
	assert.Equal(t, token.ShrAssign.String(), ">>=")
	assert.Equal(t, token.ShlAssign.String(), "<<=")
	assert.Equal(t, token.FieldAssign.String(), ":")
	assert.Equal(t, token.Invalid.String(), "")
}
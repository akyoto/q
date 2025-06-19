package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestTokenEnd(t *testing.T) {
	hello := token.Token{
		Kind:     token.Identifier,
		Position: 0,
		Length:   5,
	}

	assert.Equal(t, hello.End(), 5)
}

func TestTokenReset(t *testing.T) {
	hello := token.Token{
		Kind:     token.Identifier,
		Position: 1,
		Length:   5,
	}

	hello.Reset()
	assert.Equal(t, hello.Position, 0)
	assert.Equal(t, hello.Length, 0)
	assert.Equal(t, hello.Kind, token.Invalid)
}

func TestTokenString(t *testing.T) {
	buffer := []byte("hello, world")
	hello := token.Token{Kind: token.Identifier, Position: 0, Length: 5}
	comma := token.Token{Kind: token.Separator, Position: 5, Length: 1}
	world := token.Token{Kind: token.Identifier, Position: 7, Length: 5}

	assert.Equal(t, hello.String(buffer), "hello")
	assert.Equal(t, comma.String(buffer), ",")
	assert.Equal(t, world.String(buffer), "world")
}

func TestTokenGroups(t *testing.T) {
	assignment := token.Token{Kind: token.Assign}
	operator := token.Token{Kind: token.Add}
	keyword := token.Token{Kind: token.If}
	unary := token.Token{Kind: token.Not}
	number := token.Token{Kind: token.Number}
	comparison := token.Token{Kind: token.Equal}

	assert.True(t, assignment.IsAssignment())
	assert.True(t, operator.IsOperator())
	assert.True(t, keyword.IsKeyword())
	assert.True(t, unary.IsUnaryOperator())
	assert.True(t, number.IsNumeric())
	assert.True(t, comparison.IsComparison())
}
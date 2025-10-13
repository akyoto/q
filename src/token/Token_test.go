package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

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

func TestTokenSource(t *testing.T) {
	hello := token.Token{
		Kind:     token.Identifier,
		Position: 0,
		Length:   5,
	}

	assert.Equal(t, hello.Start(), 0)
	assert.Equal(t, hello.End(), 5)
}

func TestTokenString(t *testing.T) {
	buffer := []byte("hello, world")
	hello := token.Token{Kind: token.Identifier, Position: 0, Length: 5}
	comma := token.Token{Kind: token.Separator, Position: 5, Length: 1}
	world := token.Token{Kind: token.Identifier, Position: 7, Length: 5}

	assert.Equal(t, hello.StringFrom(buffer), "hello")
	assert.Equal(t, comma.StringFrom(buffer), ",")
	assert.Equal(t, world.StringFrom(buffer), "world")
}
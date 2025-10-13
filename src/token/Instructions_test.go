package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestInstructionsBasic(t *testing.T) {
	src := []byte("a := 1\nb := 2\n")
	tokens := token.Tokenize(src)
	nodes := []string{}

	for param := range tokens.Instructions {
		nodes = append(nodes, param.StringFrom(src))
	}

	assert.DeepEqual(t, nodes, []string{"a := 1", "b := 2"})
}

func TestInstructionsBlock(t *testing.T) {
	src := []byte("a := 1\nif x > 0 {\nx = 0\n}\nb := 2\n")
	tokens := token.Tokenize(src)
	nodes := []string{}

	for param := range tokens.Instructions {
		nodes = append(nodes, param.StringFrom(src))
	}

	assert.DeepEqual(t, nodes, []string{"a := 1", "if x > 0 {\nx = 0\n}", "b := 2"})
}

func TestInstructionsGroup(t *testing.T) {
	src := []byte("a := 1\ncall(\nx,\ny\n)\nb := 2\n")
	tokens := token.Tokenize(src)
	nodes := []string{}

	for param := range tokens.Instructions {
		nodes = append(nodes, param.StringFrom(src))
	}

	assert.DeepEqual(t, nodes, []string{"a := 1", "call(\nx,\ny\n)", "b := 2"})
}

func TestInstructionsBreak(t *testing.T) {
	src := []byte("a := 1\nb := 2\n")
	tokens := token.Tokenize(src)
	count := 0

	for range tokens.Instructions {
		if count == 1 {
			break
		}

		count++
	}
}

func TestInstructionsEOF(t *testing.T) {
	src := []byte("a := 1")
	tokens := token.Tokenize(src)
	count := 0

	for range tokens.Instructions {
		count++
	}

	assert.Equal(t, count, 1)
}

func TestInstructionsNoEOF(t *testing.T) {
	tokens := token.List{
		token.Token{Position: 0, Length: 1, Kind: token.Identifier},
	}

	count := 0

	for range tokens.Instructions {
		count++
	}

	assert.Equal(t, count, 1)
}

func TestInstructionsMultiBlock(t *testing.T) {
	src := []byte("if x == 0 { if y == 0 {} }")
	tokens := token.Tokenize(src)
	count := 0

	for range tokens.Instructions {
		count++
	}

	assert.Equal(t, count, 1)
}

func TestInstructionsMultiBlockBreak(t *testing.T) {
	src := []byte("if x == 0 { if y == 0 {} }")
	tokens := token.Tokenize(src)
	count := 0

	for range tokens.Instructions {
		count++
		break
	}

	assert.Equal(t, count, 1)
}
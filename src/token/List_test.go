package token_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestIndexKind(t *testing.T) {
	tokens := token.Tokenize([]byte("a{{}}"))
	assert.Equal(t, tokens.IndexKind(token.NewLine), -1)
	assert.Equal(t, tokens.LastIndexKind(token.NewLine), -1)
	assert.Equal(t, tokens.IndexKind(token.BlockStart), 1)
	assert.Equal(t, tokens.LastIndexKind(token.BlockStart), 2)
	assert.Equal(t, tokens.IndexKind(token.BlockEnd), 3)
	assert.Equal(t, tokens.LastIndexKind(token.BlockEnd), 4)
}

func TestSplit(t *testing.T) {
	src := []byte("1+2,3*4,5*6,7+8")
	tokens := token.Tokenize(src)
	parameters := []string{}

	for _, param := range tokens.Split {
		parameters = append(parameters, param.String(src))
	}

	assert.DeepEqual(t, parameters, []string{"1+2", "3*4", "5*6", "7+8"})
}

func TestSplitBreak(t *testing.T) {
	src := []byte("1,2")
	tokens := token.Tokenize(src)

	for range tokens.Split {
		break
	}
}

func TestSplitEmpty(t *testing.T) {
	tokens := token.List{}

	for range tokens.Split {
		t.Fail()
	}
}

func TestSplitGroups(t *testing.T) {
	src := []byte("f(1,2),g(3,4)")
	tokens := token.Tokenize(src)
	parameters := []string{}

	for _, param := range tokens.Split {
		parameters = append(parameters, param.String(src))
	}

	assert.DeepEqual(t, parameters, []string{"f(1,2)", "g(3,4)"})
}

func TestSplitGroupsTrail(t *testing.T) {
	src := []byte("f(1,2),")
	tokens := token.Tokenize(src)
	tokens = tokens[:len(tokens)-1]
	parameters := []string{}

	for _, param := range tokens.Split {
		parameters = append(parameters, param.String(src))
	}

	assert.DeepEqual(t, parameters, []string{"f(1,2)", ""})
}

func TestSplitSingle(t *testing.T) {
	src := []byte("123")
	tokens := token.Tokenize(src)
	parameters := []string{}

	for _, param := range tokens.Split {
		parameters = append(parameters, param.String(src))
	}

	assert.DeepEqual(t, parameters, []string{"123"})
}
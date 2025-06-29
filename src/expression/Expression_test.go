package expression_test

import (
	"errors"
	"testing"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestEachLeaf(t *testing.T) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	leaves := []string{}

	err := expr.EachLeaf(func(leaf *expression.Expression) error {
		leaves = append(leaves, leaf.Token.String(src))
		return nil
	})

	assert.Nil(t, err)
	assert.DeepEqual(t, leaves, []string{"1", "2", "3", "4", "5", "6", "7", "8"})

	err = expr.EachLeaf(func(leaf *expression.Expression) error {
		return errors.New("error")
	})

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}

func TestNilExpression(t *testing.T) {
	src := []byte("")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	assert.Nil(t, expr)
}

func TestNilGroup(t *testing.T) {
	src := []byte("()")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	assert.Nil(t, expr)
}

func TestRemoveChild(t *testing.T) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	left := expr.Children[0]
	right := expr.Children[1]
	expr.RemoveChild(left)
	assert.Equal(t, expr.Children[0], right)
}

func TestSource(t *testing.T) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	assert.DeepEqual(t, expr.Source, tokens)
	assert.DeepEqual(t, expr.Children[0].Source, tokens[1:8])
	assert.DeepEqual(t, expr.Children[1].Source, tokens[11:18])
	sources := []string{}

	expr.EachLeaf(func(leaf *expression.Expression) error {
		sources = append(sources, leaf.Source.String(src))
		return nil
	})

	assert.DeepEqual(t, sources, []string{"1", "2", "3", "4", "5", "6", "7", "8"})
}
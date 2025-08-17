package expression_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestLeaves(t *testing.T) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	leaves := []string{}

	for leaf := range expr.Leaves() {
		leaves = append(leaves, leaf.Token.String(src))
	}

	assert.DeepEqual(t, leaves, []string{"1", "2", "3", "4", "5", "6", "7", "8"})

	for range expr.Leaves() {
		break
	}
}

func TestLeavesBreak(t *testing.T) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)

	for range expr.Leaves() {
		break
	}
}

func TestInvalidExpression(t *testing.T) {
	src := []byte("")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	assert.Equal(t, expr.Token.Kind, token.Invalid)
}

func TestInvalidGroup(t *testing.T) {
	src := []byte("()")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	assert.Equal(t, expr.Token.Kind, token.Invalid)
}

func TestIndex(t *testing.T) {
	src := []byte("1+2")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)
	left := expr.Children[0]
	right := expr.Children[1]
	assert.Equal(t, expr.Index(left), 0)
	assert.Equal(t, expr.Index(right), 1)
	assert.Equal(t, expr.Index(expr), -1)
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

	assert.Equal(t, expr.String(src), "(+ (- (+ 1 2) (* 3 4)) (+ (- (* 5 6) 7) 8))")
	assert.Equal(t, expr.SourceString(src), string(src))

	assert.Equal(t, expr.Children[0].String(src), "(- (+ 1 2) (* 3 4))")
	assert.Equal(t, expr.Children[0].SourceString(src), "1+2-3*4")

	assert.Equal(t, expr.Children[1].String(src), "(+ (- (* 5 6) 7) 8)")
	assert.Equal(t, expr.Children[1].SourceString(src), "5*6-7+8")

	assert.Equal(t, expr.Children[0].Children[0].String(src), "(+ 1 2)")
	assert.Equal(t, expr.Children[0].Children[0].SourceString(src), "1+2")

	assert.Equal(t, expr.Children[0].Children[1].String(src), "(* 3 4)")
	assert.Equal(t, expr.Children[0].Children[1].SourceString(src), "3*4")

	assert.Equal(t, expr.Children[1].Children[0].String(src), "(- (* 5 6) 7)")
	assert.Equal(t, expr.Children[1].Children[0].SourceString(src), "5*6-7")

	assert.Equal(t, expr.Children[1].Children[0].Children[0].String(src), "(* 5 6)")
	assert.Equal(t, expr.Children[1].Children[0].Children[0].SourceString(src), "5*6")

	assert.Equal(t, expr.Children[1].Children[0].Children[1].String(src), "7")
	assert.Equal(t, expr.Children[1].Children[0].Children[1].SourceString(src), "7")

	assert.Equal(t, expr.Children[1].Children[1].String(src), "8")
	assert.Equal(t, expr.Children[1].Children[1].SourceString(src), "8")

	leaves := []string{}

	for leaf := range expr.Leaves() {
		leaves = append(leaves, leaf.SourceString(src))
	}

	assert.DeepEqual(t, leaves, []string{"1", "2", "3", "4", "5", "6", "7", "8"})
}

func TestSource2(t *testing.T) {
	src := []byte("( 1+2-(3*4) ) + ( 5*6-7+8 )")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)

	assert.Equal(t, expr.String(src), "(+ (- (+ 1 2) (* 3 4)) (+ (- (* 5 6) 7) 8))")
	assert.Equal(t, expr.SourceString(src), string(src))
}

func TestSource3(t *testing.T) {
	src := []byte("x:=2+3")
	tokens := token.Tokenize(src)
	expr := expression.Parse(tokens)

	assert.Equal(t, expr.String(src), "(:= x (+ 2 3))")
	assert.Equal(t, expr.SourceString(src), string(src))

	assert.Equal(t, expr.Children[0].String(src), "x")
	assert.Equal(t, expr.Children[0].SourceString(src), "x")

	assert.Equal(t, expr.Children[1].String(src), "(+ 2 3)")
	assert.Equal(t, expr.Children[1].SourceString(src), "2+3")
}
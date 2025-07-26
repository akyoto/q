package ast_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestAssign(t *testing.T) {
	tree, err := parse("a := 0\na = 0")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 2)

	definition := tree[0].(*ast.Define)
	assignment := tree[1].(*ast.Assign)
	assert.NotNil(t, definition.Expression)
	assert.NotNil(t, assignment.Expression)
}

func TestGroups(t *testing.T) {
	tree, err := parse("f(\nx\n)\ng(\nx\n)")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 2)

	f := tree[0].(*ast.Call)
	assert.NotNil(t, f.Expression)
	g := tree[1].(*ast.Call)
	assert.NotNil(t, g.Expression)
}

func TestIfElse(t *testing.T) {
	tree, err := parse("if x == 0 {} else {}")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 1)

	branch := tree[0].(*ast.If)
	assert.NotNil(t, branch.Condition)
	assert.NotNil(t, branch.Body)
	assert.NotNil(t, branch.Else)
}

func TestLoop(t *testing.T) {
	tree, err := parse("loop{}")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 1)
}

func TestNewLine(t *testing.T) {
	tree, err := parse("\n\n\n")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 0)
}

func TestReturn(t *testing.T) {
	tree, err := parse("return")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 1)

	ret := tree[0].(*ast.Return)
	assert.Nil(t, ret.Values)
}

func TestReturnValues(t *testing.T) {
	tree, err := parse("return 42")
	assert.Nil(t, err)
	assert.Equal(t, len(tree), 1)

	ret := tree[0].(*ast.Return)
	assert.Equal(t, len(ret.Values), 1)
	assert.Equal(t, ret.Values[0].Token.Kind, token.Number)
}

func parse(code string) (ast.AST, error) {
	src := []byte(code)
	tokens := token.Tokenize(src)
	file := &fs.File{Bytes: src, Tokens: tokens}
	return ast.Parse(tokens, file)
}
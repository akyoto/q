package expression

import (
	"strings"

	"git.urbach.dev/cli/q/src/token"
)

// Expression is a tree that can represent a mathematical expression with precedence levels.
type Expression struct {
	Parent     *Expression
	Children   []*Expression
	Token      token.Token
	Source     token.List
	precedence int8
}

// AddChild adds a child to the expression.
func (expr *Expression) AddChild(child *Expression) {
	if expr.Children == nil {
		expr.Children = make([]*Expression, 0, 2)
	}

	expr.Children = append(expr.Children, child)
	child.Parent = expr
}

// Reset resets all values to the default.
func (expr *Expression) Reset() {
	expr.Parent = nil

	if expr.Children != nil {
		expr.Children = expr.Children[:0]
	}

	expr.Token.Reset()
	expr.precedence = 0
}

// EachLeaf iterates through all leaves in the tree.
func (expr *Expression) EachLeaf(call func(*Expression) error) error {
	if expr.IsLeaf() {
		return call(expr)
	}

	for _, child := range expr.Children {
		err := child.EachLeaf(call)

		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveChild removes a child from the expression.
func (expr *Expression) RemoveChild(child *Expression) {
	for i, c := range expr.Children {
		if c == child {
			expr.Children = append(expr.Children[:i], expr.Children[i+1:]...)
			child.Parent = nil
			return
		}
	}
}

// InsertAbove replaces this expression in its parent's children with the given new parent,
// and attaches this expression as a child of the new parent. Effectively, it promotes the
// given tree above the current node. It assumes that the caller is the last child.
func (expr *Expression) InsertAbove(tree *Expression) {
	if expr.Parent != nil {
		expr.Parent.Children[len(expr.Parent.Children)-1] = tree
		tree.Parent = expr.Parent
	}

	tree.AddChild(expr)
}

// IsLeaf returns true if the expression has no children.
func (expr *Expression) IsLeaf() bool {
	return len(expr.Children) == 0
}

// LastChild returns the last child.
func (expr *Expression) LastChild() *Expression {
	return expr.Children[len(expr.Children)-1]
}

// String generates a textual representation of the expression.
func (expr *Expression) String(source []byte) string {
	builder := strings.Builder{}
	expr.write(&builder, source)
	return builder.String()
}
package expression

import (
	"strings"

	"github.com/akyoto/q/build/token"
)

// Expression is a binary tree with an operator on each node.
type Expression struct {
	Value    token.Token
	Children []*Expression
	Parent   *Expression
}

// AddChild adds a child to the expression.
func (expr *Expression) AddChild(t token.Token) {
	if expr.Value.Kind == token.Invalid {
		expr.Value = t
		return
	}

	expr.Children = append(expr.Children, &Expression{
		Value:  t,
		Parent: expr,
	})
}

// LastChild returns the last child.
func (expr *Expression) LastChild() *Expression {
	return expr.Children[len(expr.Children)-1]
}

// IsLeaf returns true if the expression is a leaf node with no children.
func (expr *Expression) IsLeaf() bool {
	return len(expr.Children) == 0
}

// String generates a textual representation of the expression.
func (expr *Expression) String() string {
	builder := strings.Builder{}
	expr.write(&builder)
	return builder.String()
}

// write generates a textual representation of the expression.
func (expr *Expression) write(builder *strings.Builder) {
	if len(expr.Children) == 0 {
		builder.WriteString(expr.Value.Text())
		return
	}

	builder.WriteByte('(')

	for index, operand := range expr.Children {
		operand.write(builder)

		if index != len(expr.Children)-1 {
			builder.WriteString(expr.Value.Text())
		}
	}

	builder.WriteByte(')')
}

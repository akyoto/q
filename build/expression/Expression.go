package expression

import (
	"bytes"
	"strings"

	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

var functionCallToken = token.Token{Kind: token.Operator, Bytes: []byte{'(', ')'}}

// Expression is a binary tree with an operator on each node.
type Expression struct {
	Token    token.Token
	Children []*Expression
	Parent   *Expression
	Register *register.Register
	Grouped  bool
}

// New creates a new expression.
func New() *Expression {
	return pool.Get().(*Expression)
}

// AddChild adds a child to the expression.
func (expr *Expression) AddChild(operand *Expression) {
	if operand.Parent != nil {
		operand.Parent.RemoveChild(operand)
	}

	operand.Parent = expr
	expr.Children = append(expr.Children, operand)
}

// PrependChild adds a child to the expression at the start.
func (expr *Expression) PrependChild(operand *Expression) {
	if operand.Parent != nil {
		operand.Parent.RemoveChild(operand)
	}

	operand.Parent = expr
	expr.Children = append([]*Expression{operand}, expr.Children...)
}

// RemoveChild removes a child from the expression.
func (expr *Expression) RemoveChild(operand *Expression) {
	for i, child := range expr.Children {
		if child == operand {
			expr.Children = append(expr.Children[:i], expr.Children[i+1:]...)
			return
		}
	}
}

// AddToken adds a token to the expression.
func (expr *Expression) AddToken(t token.Token) *Expression {
	child := New()
	child.Token = t
	child.Parent = expr
	expr.Children = append(expr.Children, child)
	return child
}

// EachOperation iterates the operations in the tree.
func (expr *Expression) EachOperation(callBack func(*Expression) error) error {
	if expr.IsLeaf() {
		return nil
	}

	for _, child := range expr.Children {
		if child.IsLeaf() {
			continue
		}

		err := child.EachOperation(callBack)

		if err != nil {
			return err
		}
	}

	return callBack(expr)
}

// SortByRegisterCount sorts the children by register count.
func (expr *Expression) SortByRegisterCount() {
	if expr.IsLeaf() {
		return
	}

	for _, child := range expr.Children {
		child.SortByRegisterCount()
	}

	if expr.Token.Kind == token.Operator && spec.Operators[string(expr.Token.Bytes)].OperandOrderImportant {
		return
	}

	left := expr.Children[0]
	right := expr.Children[1]

	leftCount := left.RegisterCount()
	rightCount := right.RegisterCount()

	if rightCount > leftCount {
		expr.Children[0] = right
		expr.Children[1] = left
	}
}

// RegisterCount returns the number of registers
// needed to calculate this expression tree.
func (expr *Expression) RegisterCount() int {
	count := 0

	for _, child := range expr.Children {
		count += child.RegisterCount()
	}

	if expr.Token.Kind == token.Operator && count == 0 {
		count = 1
	}

	return count
}

// LastChild returns the last child.
func (expr *Expression) LastChild() *Expression {
	return expr.Children[len(expr.Children)-1]
}

// IsLeaf returns true if the expression is a leaf node with no children.
func (expr *Expression) IsLeaf() bool {
	return len(expr.Children) == 0
}

// IsFunctionCall returns true if the expression is a function call.
func (expr *Expression) IsFunctionCall() bool {
	return expr.Token.Kind == functionCallToken.Kind && bytes.Equal(expr.Token.Bytes, functionCallToken.Bytes)
}

// Close puts the expression back into the memory pool.
func (expr *Expression) Close() {
	for _, child := range expr.Children {
		child.Close()
	}

	expr.Token.Kind = token.Invalid
	expr.Children = expr.Children[:0]
	expr.Parent = nil
	expr.Register = nil
	pool.Put(expr)
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
		builder.WriteString(expr.Token.Text())
		return
	}

	children := expr.Children
	operator := expr.Token.Text()

	if expr.IsFunctionCall() {
		builder.WriteString(children[0].Token.Text())
		children = children[1:]
		operator = ","
	}

	builder.WriteByte('(')

	for index, operand := range children {
		operand.write(builder)

		if index != len(children)-1 || (len(children) == 1 && !expr.IsFunctionCall()) {
			builder.WriteString(operator)
		}
	}

	builder.WriteByte(')')
}

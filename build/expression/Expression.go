package expression

import (
	"strings"

	"github.com/akyoto/q/build/operators"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/q/build/types"
)

// Expression is a binary tree with an operator on each node.
type Expression struct {
	Token          token.Token
	Children       []*Expression
	Parent         *Expression
	Register       *register.Register
	Type           *types.Type
	IsFunctionCall bool
}

// New creates a new expression.
func New() *Expression {
	return pool.Get().(*Expression)
}

// AddChild adds a child to the expression.
func (expr *Expression) AddChild(child *Expression) {
	if child.Parent != nil {
		child.Parent.RemoveChild(child)
	}

	child.Parent = expr
	expr.Children = append(expr.Children, child)
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

// SetParent sets the parent of the expression.
func (expr *Expression) SetParent(parent *Expression) {
	parent.AddChild(expr)
}

// IsEmpty tells you whether the expression is empty or not.
func (expr *Expression) IsEmpty() bool {
	return expr.Token.Kind == token.Invalid && len(expr.Children) == 0
}

// EachOperation iterates the operations in the tree.
func (expr *Expression) EachOperation(callBack func(*Expression) error) error {
	if expr.IsLeaf() {
		return nil
	}

	// Don't descend into the parameters of function calls.
	// We rely on the compiler using CallExpression for each of them.
	if expr.IsFunctionCall {
		return callBack(expr)
	}

	for _, child := range expr.Children {
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

	if expr.IsFunctionCall || (expr.Token.Kind == token.Operator && operators.All[string(expr.Token.Bytes)].OperandOrderImportant) {
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

// Replace replaces the expression with the given expression.
func (expr *Expression) Replace(other *Expression) {
	parent := expr.Parent
	*expr = *other
	expr.Parent = parent
}

// IsLeaf returns true if the expression is a leaf node with no children.
func (expr *Expression) IsLeaf() bool {
	return !expr.IsFunctionCall && len(expr.Children) == 0
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
	expr.Type = nil
	expr.IsFunctionCall = false
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
	if !expr.IsFunctionCall && len(expr.Children) == 0 {
		builder.WriteString(expr.Token.Text())
		return
	}

	children := expr.Children
	operator := expr.Token.Text()

	if expr.IsFunctionCall {
		builder.WriteString(expr.Token.Text())
		operator = ","
	}

	builder.WriteByte('(')

	for index, operand := range children {
		operand.write(builder)

		if index != len(children)-1 || (len(children) == 1 && !expr.IsFunctionCall) {
			builder.WriteString(operator)
		}
	}

	builder.WriteByte(')')
}

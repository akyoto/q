package expression

import (
	"strings"

	"github.com/akyoto/q/build/spec"
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

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	current := &Expression{}
	stack := []*Expression{current}
	goUp := false

	for index, t := range tokens {
		switch t.Kind {
		case token.Identifier, token.Number, token.Text:
			current.AddChild(t)

			// In case an operator priority was enforced,
			// we need to go back up to the original node.
			if goUp {
				current = current.Parent
				goUp = false
			}

		case token.GroupStart:
			group := &Expression{
				Parent: current,
			}

			current = group
			stack = append(stack, group)

		case token.GroupEnd:
			if len(current.Children) == 0 {
				current.Parent.AddChild(current.Value)
			} else {
				if len(current.Parent.Children) == 0 {
					current.Parent.Value = current.Value
					current.Parent.Children = current.Children
				} else {
					current.Parent.Children = append(current.Parent.Children, current)
				}
			}

			stack = stack[:len(stack)-1]
			current = stack[len(stack)-1]

		case token.Operator:
			// Turn identifier into an operation
			if current.IsLeaf() {
				current.Children = append(current.Children, &Expression{
					Value:  current.Value,
					Parent: current,
				})

				current.Value = t
				continue
			}

			// Calculate priority
			if index > 0 && tokens[index-1].Kind != token.GroupEnd && len(current.Children) >= 2 && current.LastChild().Value.Kind != token.Operator {
				priority := spec.Operators[t.Text()]
				lastPriority := spec.Operators[current.Value.Text()]

				if priority > lastPriority {
					// Expression: 1 + 2 * 3
					//                 ^
					//                 lastChild
					//                 ^^^
					//                 subExpression
					lastChild := current.Children[len(current.Children)-1]

					subExpression := &Expression{
						Value:    t,
						Children: []*Expression{lastChild},
						Parent:   current,
					}

					current.Children[len(current.Children)-1] = subExpression
					current = subExpression
					goUp = true
					continue
				}
			}

			newOperator := &Expression{
				Value:    t,
				Children: []*Expression{current},
				Parent:   current.Parent,
			}

			current.Parent = newOperator
			current = newOperator
			stack[len(stack)-1] = newOperator
		}
	}

	return stack[0], nil
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

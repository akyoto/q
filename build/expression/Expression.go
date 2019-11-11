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

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	current := &Expression{}

	for _, t := range tokens {
		switch t.Kind {
		case token.Identifier, token.Number, token.Text:
			// Identity
			if current.Value.Kind == token.Invalid {
				current.Value = t
				continue
			}

			// Add operand
			current.Children = append(current.Children, &Expression{
				Value:  t,
				Parent: current,
			})

		case token.GroupStart:
			group := &Expression{
				Parent: current,
			}

			current.Children = append(current.Children, group)
			current = group

		case token.GroupEnd:
			// Identity group
			if current.Value.Kind != token.Operator && current.Parent.Value.Kind == token.Invalid {
				current.Parent.Value = current.Value
				current.Parent.Children = current.Parent.Children[:len(current.Parent.Children)-1]
			}

			// Same operator as parent?
			if current.Value.Kind == token.Operator && current.Parent.Value.Kind == token.Operator && current.Value.Text() == current.Parent.Value.Text() {
				for _, child := range current.Children {
					child.Parent = current.Parent
				}

				current.Parent.Children = current.Parent.Children[:len(current.Parent.Children)-1]
				current.Parent.Children = append(current.Parent.Children, current.Children...)
			}

			current = current.Parent

		case token.Operator:
			// Same operator
			if current.Value.Kind == token.Operator && current.Value.Text() == t.Text() {
				continue
			}

			// Same operator as parent
			if current.Parent != nil && current.Value.Kind == token.Operator && current.Parent.Value.Kind == token.Operator && current.Parent.Value.Text() == t.Text() {
				current = current.Parent
				continue
			}

			// Turn identifier into an operation
			if current.Value.Kind == token.Identifier || current.Value.Kind == token.Number || current.Value.Kind == token.Text {
				current.Children = append(current.Children, &Expression{
					Value:  current.Value,
					Parent: current,
				})

				current.Value = t
				continue
			}

			// If it started with a group, we need to set the operator
			if current.Value.Kind == token.Invalid {
				current.Value = t
				continue
			}

			// Calculate precedence
			precedence := spec.Operators[t.Text()]
			currentPrecedence := spec.Operators[current.Value.Text()]

			if precedence > currentPrecedence {
				lastChild := current.Children[len(current.Children)-1]

				lastChild = &Expression{
					Value:    t,
					Children: []*Expression{lastChild},
					Parent:   current,
				}

				current.Children[len(current.Children)-1] = lastChild
				current = lastChild
			} else {
				// Current expression becomes a child of right expression
				right := &Expression{
					Value:    t,
					Children: []*Expression{current},
					Parent:   current.Parent,
				}

				if current.Parent != nil {
					current.Parent.Children[len(current.Parent.Children)-1] = right
				}

				current.Parent = right
				current = right
			}
		}
	}

	for current.Parent != nil {
		current = current.Parent
	}

	return current, nil
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

	if expr.Parent != nil {
		builder.WriteByte('(')
	}

	for index, operand := range expr.Children {
		operand.write(builder)

		if index != len(expr.Children)-1 {
			builder.WriteString(expr.Value.Text())
		}
	}

	if expr.Parent != nil {
		builder.WriteByte(')')
	}
}

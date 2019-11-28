package expression

import (
	"fmt"

	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	if len(tokens) == 1 {
		return FromToken(tokens[0]), nil
	}

	// Save the number of parentheses we encounter so we can
	// notice the end of a new group when group level is 0.
	groupLevel := 0
	groupPosition := 0

	// Create a root node and use it as our current expression.
	current := New()

	// Next current is used for when we want to change the current node
	// after we receive the next operand.
	var nextCurrent *Expression

	// We iterate over all tokens and adjust the expression tree as we go.
	for i, t := range tokens {
		fmt.Printf("[%s] %v\n", t, current)

		switch t.Kind {
		case token.GroupStart:
			if groupLevel == 0 {
				groupPosition = i + 1
			}

			groupLevel++
			continue

		case token.GroupEnd:
			groupLevel--

			if groupLevel == 0 {
				operand, err := FromTokens(tokens[groupPosition:i])

				if err != nil {
					return nil, err
				}

				current.AddChild(operand)

				if nextCurrent != nil {
					current = nextCurrent
					nextCurrent = nil
				}
			}

			continue

		default:
			if groupLevel != 0 {
				continue
			}
		}

		switch t.Kind {
		case token.Identifier, token.Number, token.Text:
			operand := FromToken(t)
			current.AddChild(operand)

			if nextCurrent != nil {
				current = nextCurrent
				nextCurrent = nil
			}

		case token.Operator:
			if current.Token.Kind != token.Operator {
				current.Token = t
				continue
			}

			// Compare operator priority
			oldOperator := current.Token.Text()
			oldOperatorPriority := spec.Operators[oldOperator].Priority
			newOperator := t.Text()
			newOperatorPriority := spec.Operators[newOperator].Priority

			if newOperatorPriority > oldOperatorPriority {
				// Let's say we have the expression (1 + 2 * 3)
				// At first, we encountered 1 + 2 and generated this tree:
				//   + (current)
				//  / \
				// 1   2

				// Now we encountered a higher priority operator.
				// We need to take the last operand and replace it with the higher priority operation:
				//   +
				//  / \
				// 1   * (current)
				//    /
				//   2

				// Take the last child
				lastChild := current.LastChild()

				// Create a new expression for the higher priority operation
				newOperation := New()
				newOperation.Token = t
				newOperation.AddChild(lastChild)
				newOperation.SetParent(current)

				// The new operator becomes the current expression until we added the second operand entirely.
				nextCurrent = current
				current = newOperation
				continue
			}

			newOperation := New()
			newOperation.Token = t
			current.SetParent(newOperation)
			current = newOperation
		}
	}

	fmt.Printf(" =  %v\n", current)

	// Walk up the tree and return the top level node.
	for current.Parent != nil {
		current = current.Parent
	}

	// If we only have 1 child in an invalid operation,
	// replace the result with the child itself.
	// This turns expressions like (123) into 123.
	if current.Token.Kind == token.Invalid && len(current.Children) == 1 {
		current = current.Children[0]
		current.Parent.Children = nil
		current.Parent = nil
	}

	return current, nil
}

// FromToken generates an expression for a single token.
func FromToken(t token.Token) *Expression {
	operand := New()
	operand.Token = t
	return operand
}

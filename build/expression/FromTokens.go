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

	fmt.Println("---")
	groupLevel := 0
	groupStart := 0
	root := New()
	current := root

	onOperator := func(t token.Token) {
		if current.Token.Kind != token.Operator {
			current.Token = t
			return
		}

		// Compare operator priority
		oldOperator := current.Token.Text()
		oldOperatorPriority := spec.Operators[oldOperator].Priority
		newOperator := t.Text()
		newOperatorPriority := spec.Operators[newOperator].Priority

		if newOperatorPriority > oldOperatorPriority {
			fmt.Println("HIGHER")
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
			current = newOperation
			return
		}

		newOperation := New()
		newOperation.Token = t
		current.SetParent(newOperation)
		current = newOperation
		root = newOperation
	}

	onOperand := func(operand *Expression) {
		current.AddChild(operand)
	}

	for i, t := range tokens {
		switch t.Kind {
		case token.GroupStart:
			if groupLevel == 0 {
				groupStart = i + 1
			}

			groupLevel++
			continue

		case token.GroupEnd:
			groupLevel--

			if groupLevel == 0 {
				operand, err := FromTokens(tokens[groupStart:i])

				if err != nil {
					return nil, err
				}

				onOperand(operand)
			}

			continue

		default:
			if groupLevel != 0 {
				continue
			}
		}

		switch t.Kind {
		case token.Identifier:
			operand := FromToken(t)
			onOperand(operand)

		case token.Number, token.Text:
			operand := FromToken(t)
			onOperand(operand)

		case token.Operator:
			onOperator(t)
		}
	}

	fmt.Println("---")
	return root, nil
}

// FromToken generates an expression for a single token.
func FromToken(t token.Token) *Expression {
	operand := New()
	operand.Token = t
	return operand
}

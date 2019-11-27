package expression

import (
	"fmt"

	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	fmt.Println("---", tokens, "---")
	operands := []*Expression{}
	operators := []token.Token{}
	groupLevel := 0
	groupStart := 0

	for i, t := range tokens {
		switch t.Kind {
		case token.GroupStart:
			groupLevel++

			if groupLevel == 1 {
				groupStart = i + 1
			}

		case token.GroupEnd:
			groupLevel--

			if groupLevel == 0 {
				operand, err := FromTokens(tokens[groupStart:i])

				if err != nil {
					return nil, err
				}

				operands = append(operands, operand)
			}
		}

		if groupLevel != 0 {
			continue
		}

		switch t.Kind {
		case token.Identifier:
			operand := New()
			operand.Token = t
			operands = append(operands, operand)

		case token.Number, token.Text:
			operand := New()
			operand.Token = t
			operands = append(operands, operand)

		case token.Operator:
			operators = append(operators, t)
		}
	}

	root := New()

	if len(operands) == 0 {
		return root, nil
	}

	if len(operators) == 0 {
		return operands[0], nil
	}

	current := root
	var lastOperator []byte

	for i := len(operators) - 1; i >= 0; i-- {
		operator := operators[i]
		right := operands[i+1]

		if lastOperator != nil {
			rightOperatorPriority := spec.Operators[string(lastOperator)].Priority
			leftOperatorPriority := spec.Operators[string(operator.Bytes)].Priority

			if rightOperatorPriority > leftOperatorPriority {
				parent := current.Parent

				*current = *right

				fmt.Println("current =", current)
				fmt.Println("parent =", parent)
				fmt.Println("root =", root)

				current = New()

				if root == parent {
					root = current
					fmt.Println("NEW ROOT", root)
				} else if parent.Parent != nil {
					parent.Parent.PrependChild(current)
				}

				right = parent
			}
		}

		current.Token = operator
		var left *Expression

		if i == 0 {
			left = operands[0]
		} else {
			left = New()
		}

		fmt.Println("current =", current, "in", current.Parent)
		fmt.Println("left =", left)
		fmt.Println("right =", right)
		fmt.Println("operator =", operator)

		current.AddChild(left)
		current.AddChild(right)
		current = left
		lastOperator = operator.Bytes
	}

	return root, nil
}

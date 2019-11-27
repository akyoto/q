package expression

import (
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	current := New()
	stack := []*Expression{current}
	var goUp *Expression

	for index, t := range tokens {
		switch t.Kind {
		case token.Identifier:
			// Function calls
			if index != len(tokens)-1 && tokens[index+1].Kind == token.GroupStart {
				call := current.AddToken(functionCallToken)
				current = call
			}

			current.AddToken(t)

		case token.Number:
			current.AddToken(t)

		case token.Text:
			current.AddToken(t)

		case token.GroupStart:
			if current.IsFunctionCall() {
				stack = append(stack, current)
				continue
			}

			// Create a group and assign the parent.
			// We don't add it to the parent yet.
			// This will happen at the end of the group.
			group := New()
			group.Parent = current

			current = group
			stack = append(stack, group)

		case token.GroupEnd:
			group := stack[len(stack)-1]

			if group.IsFunctionCall() {
				stack = stack[:len(stack)-1]
				current = stack[len(stack)-1]
				continue
			}

			group.Grouped = true

			// When we encounter the end of a group,
			// we add it to the parent.
			if group.Parent != nil {
				if group.IsLeaf() {
					// If the group contains no operations,
					// add the group token to the parent.
					// Empty groups will automatically be filtered out.
					group.Parent.AddToken(group.Token)
				} else {
					// The group contains operations.
					if group.Parent.IsLeaf() {
						*group.Parent = *group
					} else {
						group.Parent.Children = append(group.Parent.Children, group)
					}
				}
			}

			stack = stack[:len(stack)-1]
			current = stack[len(stack)-1]

		case token.Separator:
			goUp = nil

			for !current.IsFunctionCall() {
				current = current.Parent
			}

		case token.Operator:
			// Turn identifier into an operation
			if current.IsLeaf() {
				child := New()
				child.Token = current.Token
				child.Parent = current

				current.Children = append(current.Children, child)
				current.Token = t
				continue
			}

			// Change last child from single token to operation
			if !current.Grouped { // && len(current.Children) >= 1 && current.LastChild().Token.Kind != token.Operator
				// Calculate priority
				priority := spec.Operators[string(t.Bytes)].Priority
				lastPriority := 0
				isCall := current.IsFunctionCall()

				if isCall {

				} else {
					lastPriority = spec.Operators[string(current.Token.Bytes)].Priority
				}

				if priority > lastPriority {
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

					parent := current

					// Take the last child
					lastChild := parent.LastChild()

					// Create a new expression for the higher priority operation
					newOperator := New()
					newOperator.Token = t
					newOperator.Children = []*Expression{lastChild}
					newOperator.Parent = parent

					// Replace the last child of the parent
					parent.Children[len(parent.Children)-1] = newOperator

					// The new operator becomes the current expression until we added the second operand entirely.
					current = newOperator

					if !isCall {
						goUp = newOperator
					}

					continue
				}
			}

			// current becomes a child of the new operation
			newOperator := New()
			newOperator.Token = t
			newOperator.Children = []*Expression{current}
			newOperator.Parent = current.Parent

			current.Parent = newOperator
			current = newOperator
			stack[len(stack)-1] = newOperator
		}

		// In case an operator priority was enforced,
		// we need to go back up to the original node.
		if goUp != nil && goUp == current {
			current = current.Parent
			goUp = nil
		}
	}

	return stack[0], nil
}

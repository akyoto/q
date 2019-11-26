package expression

import (
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// FromTokens generates an expression tree from tokens.
func FromTokens(tokens []token.Token) (*Expression, error) {
	current := New()
	stack := []*Expression{current}
	goUp := false

	for index, t := range tokens {
		switch t.Kind {
		case token.Identifier:
			// Function calls
			if index != len(tokens)-1 && tokens[index+1].Kind == token.GroupStart {
				call := current.AddChild(token.Token{Kind: token.Operator, Bytes: nil})
				call.AddChild(t)
				current = call
				continue
			}

			current.AddChild(t)

			// In case an operator priority was enforced,
			// we need to go back up to the original node.
			if goUp {
				current = current.Parent
				goUp = false
			}

		case token.Number, token.Text:
			current.AddChild(t)

			// In case an operator priority was enforced,
			// we need to go back up to the original node.
			if goUp {
				current = current.Parent
				goUp = false
			}

		case token.GroupStart:
			group := New()
			group.Parent = current

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

		case token.Operator, token.Separator:
			// Turn identifier into an operation
			if current.IsLeaf() {
				child := New()
				child.Value = current.Value
				child.Parent = current

				current.Children = append(current.Children, child)
				current.Value = t
				continue
			}

			// Calculate priority
			if index > 0 && tokens[index-1].Kind != token.GroupEnd && len(current.Children) >= 2 && current.LastChild().Value.Kind != token.Operator {
				priority := spec.Operators[string(t.Bytes)].Priority
				lastPriority := spec.Operators[string(current.Value.Bytes)].Priority

				if priority > lastPriority {
					// Expression: 1 + 2 * 3
					//                 ^
					//                 lastChild
					//                 ^^^
					//                 subExpression
					lastChild := current.Children[len(current.Children)-1]

					subExpression := New()
					subExpression.Value = t
					subExpression.Children = []*Expression{lastChild}
					subExpression.Parent = current

					current.Children[len(current.Children)-1] = subExpression
					current = subExpression
					goUp = true
					continue
				}
			}

			newOperator := New()
			newOperator.Value = t
			newOperator.Children = []*Expression{current}
			newOperator.Parent = current.Parent

			current.Parent = newOperator
			current = newOperator
			stack[len(stack)-1] = newOperator
		}
	}

	return stack[0], nil
}

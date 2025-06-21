package expression

import (
	"math"

	"git.urbach.dev/cli/q/src/token"
)

// Parse generates an expression tree from tokens.
func Parse(tokens []token.Token) *Expression {
	var (
		cursor        *Expression
		root          *Expression
		groupLevel    = 0
		groupPosition = 0
	)

	for i, t := range tokens {
		if t.Kind == token.GroupStart || t.Kind == token.ArrayStart {
			groupLevel++

			if groupLevel == 1 {
				groupPosition = i + 1
			}

			continue
		}

		if t.Kind == token.GroupEnd || t.Kind == token.ArrayEnd {
			groupLevel--

			if groupLevel != 0 {
				continue
			}

			// Function call or array access
			if isComplete(cursor) {
				parameters := NewList(tokens[groupPosition:i])
				node := New()
				node.Token.Position = tokens[groupPosition].Position

				switch t.Kind {
				case token.GroupEnd:
					node.Token.Kind = token.Call
				case token.ArrayEnd:
					node.Token.Kind = token.Array
				}

				node.precedence = precedence(node.Token.Kind)

				if cursor.Token.IsOperator() && node.precedence > cursor.precedence {
					cursor.LastChild().InsertAbove(node)
				} else {
					if cursor == root {
						root = node
					}

					cursor.InsertAbove(node)
				}

				for _, param := range parameters {
					node.AddChild(param)
				}

				cursor = node
				continue
			}

			group := Parse(tokens[groupPosition:i])

			if group == nil {
				continue
			}

			group.precedence = math.MaxInt8

			if cursor == nil {
				if t.Kind == token.ArrayEnd {
					cursor = New()
					cursor.Token.Position = tokens[groupPosition].Position
					cursor.Token.Kind = token.Array
					cursor.precedence = precedence(token.Array)
					cursor.AddChild(group)
					root = cursor
				} else {
					cursor = group
					root = group
				}
			} else {
				cursor.AddChild(group)
			}

			continue
		}

		if groupLevel > 0 {
			continue
		}

		if t.Kind == token.Identifier || t.Kind == token.Number || t.Kind == token.String || t.Kind == token.Rune {
			if cursor != nil {
				node := NewLeaf(t)
				cursor.AddChild(node)
			} else {
				cursor = NewLeaf(t)
				root = cursor
			}

			continue
		}

		if !t.IsOperator() {
			continue
		}

		if cursor == nil {
			cursor = NewLeaf(t)
			cursor.precedence = precedence(t.Kind)
			root = cursor
			continue
		}

		node := NewLeaf(t)
		node.precedence = precedence(t.Kind)

		if cursor.Token.IsOperator() {
			oldPrecedence := cursor.precedence
			newPrecedence := node.precedence

			if newPrecedence > oldPrecedence {
				if len(cursor.Children) == numOperands(cursor.Token.Kind) {
					cursor.LastChild().InsertAbove(node)
				} else {
					cursor.AddChild(node)
				}
			} else {
				start := cursor

				for start != nil {
					precedence := start.precedence

					if precedence < newPrecedence {
						start.LastChild().InsertAbove(node)
						break
					}

					if precedence == newPrecedence {
						if start == root {
							root = node
						}

						start.InsertAbove(node)
						break
					}

					start = start.Parent
				}

				if start == nil {
					root.InsertAbove(node)
					root = node
				}
			}
		} else {
			node.AddChild(cursor)
			root = node
		}

		cursor = node
	}

	return root
}
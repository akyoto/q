package expression

import (
	"math"

	"git.urbach.dev/cli/q/src/token"
)

// handleGroupEnd handles anything that uses a group like a function call or an array access.
func handleGroupEnd(tokens token.List, root *Expression, cursor *Expression, groupPosition int, i int, t token.Token) (*Expression, *Expression) {
	if isComplete(cursor) {
		node := New()

		switch t.Kind {
		case token.GroupEnd:
			node.Token.Kind = token.Call
		case token.ArrayEnd:
			node.Token.Kind = token.Array
		case token.BlockEnd:
			node.Token.Kind = token.Struct
		}

		node.precedence = precedence(node.Token.Kind)
		node.Token.Position = tokens[groupPosition].Position
		identifier := cursor

		if cursor.Token.Kind.IsOperator() && node.precedence > cursor.precedence {
			identifier = cursor.LastChild()
		} else if cursor == root {
			root = node
		}

		identifier.InsertAbove(node)

		if identifier.Token.Kind == token.New {
			node.AddChild(&newTypeExpression(tokens[groupPosition:i]).Expression)
		} else {
			parameters := NewList(tokens[groupPosition:i])

			for _, param := range parameters {
				node.AddChild(param)
			}
		}

		cursor = node
		return root, cursor
	}

	group := Parse(tokens[groupPosition:i])
	group.precedence = math.MaxInt8

	if group.Token.Kind == token.Invalid {
		group.Token.Position = tokens[groupPosition].Position
	}

	if t.Kind == token.ArrayEnd {
		array := New()
		array.Token.Position = tokens[groupPosition].Position
		array.Token.Kind = token.Array
		array.precedence = precedence(token.Array)

		if cursor == nil {
			cursor = array
			cursor.AddChild(group)
			root = cursor
		} else {
			array.AddChild(group)
			cursor.AddChild(array)
		}

		return root, cursor
	}

	if cursor == nil {
		cursor = group
		root = group
	} else {
		cursor.AddChild(group)
	}

	return root, cursor
}
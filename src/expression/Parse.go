package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// Parse generates an expression tree from tokens.
func Parse(tokens token.List) *Expression {
	var (
		cursor *Expression
		root   *Expression
		i      uint
	)

loop:
	for i < uint(len(tokens)) {
		t := tokens[i]

		switch t.Kind {
		case token.GroupStart, token.ArrayStart, token.BlockStart:
			i++
			groupLevel := 1
			groupPosition := i

			for i < uint(len(tokens)) {
				t = tokens[i]

				switch t.Kind {
				case token.GroupStart, token.ArrayStart, token.BlockStart:
					groupLevel++
				case token.GroupEnd, token.ArrayEnd, token.BlockEnd:
					groupLevel--

					if groupLevel == 0 {
						root, cursor = handleGroupEnd(tokens, root, cursor, groupPosition, i, t)
						i++
						continue loop
					}
				}

				i++
			}

			break loop
		}

		switch {
		case cursor != nil && cursor.Token.Kind == token.Cast && len(cursor.Children) < 2:
			cursor.AddChild(&newTypeExpression(tokens[i:]).Expression)
			return root

		case t.Kind.IsLiteral():
			root, cursor = handleLiteral(root, cursor, t)

		case !t.Kind.IsOperator():
			// do nothing

		case cursor == nil:
			cursor = newLeaf(t)
			cursor.precedence = precedence(t.Kind)
			root = cursor

		default:
			node := newLeaf(t)
			node.precedence = precedence(t.Kind)

			if cursor.Token.Kind.IsOperator() {
				root = handleOperator(root, cursor, node)
			} else {
				node.AddChild(cursor)
				root = node
			}

			cursor = node
		}

		i++
	}

	if root == nil {
		root = New()
	}

	return root
}
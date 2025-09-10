package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// Parse generates an expression tree from tokens.
func Parse(tokens token.List) *Expression {
	var (
		cursor        *Expression
		root          *Expression
		groupLevel    = 0
		groupPosition = 0
	)

	for i, t := range tokens {
		if t.Kind == token.GroupStart || t.Kind == token.ArrayStart || t.Kind == token.BlockStart {
			groupLevel++

			if groupLevel == 1 {
				groupPosition = i + 1
			}

			continue
		}

		if t.Kind == token.GroupEnd || t.Kind == token.ArrayEnd || t.Kind == token.BlockEnd {
			groupLevel--

			if groupLevel != 0 {
				continue
			}

			root, cursor = handleGroupEnd(tokens, root, cursor, groupPosition, i, t)
			continue
		}

		if groupLevel > 0 {
			continue
		}

		if cursor != nil && cursor.Token.Kind == token.Cast && len(cursor.Children) < 2 {
			cursor.AddChild(&newTypeExpression(tokens[i:]).Expression)
			return root
		}

		if t.Kind.IsLiteral() {
			root, cursor = handleLiteral(root, cursor, t)
			continue
		}

		if !t.Kind.IsOperator() {
			continue
		}

		if cursor == nil {
			cursor = newLeaf(t)
			cursor.precedence = precedence(t.Kind)
			root = cursor
			continue
		}

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

	if root == nil {
		root = New()
	}

	return root
}
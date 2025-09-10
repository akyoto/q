package expression

import "git.urbach.dev/cli/q/src/token"

// handleLiteral either adds a leaf to the cursor or becomes the cursor.
func handleLiteral(root *Expression, cursor *Expression, t token.Token) (*Expression, *Expression) {
	if cursor != nil {
		node := newLeaf(t)

		if cursor.Token.Kind == token.Range && len(cursor.Children) == 0 {
			cursor.AddChild(New())
		}

		cursor.AddChild(node)
	} else {
		cursor = newLeaf(t)
		root = cursor
	}

	return root, cursor
}
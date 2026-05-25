package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// handleLiteral either adds a leaf to the cursor or becomes the cursor.
func handleLiteral(root *Expression, cursor *Expression, t token.Token) (*Expression, *Expression) {
	if cursor != nil {
		node := newLeaf(t)

		if len(cursor.Children) == 0 {
			switch cursor.Token.Kind {
			case token.Range, token.Dot:
				invalid := New()
				invalid.Token.Position = cursor.Token.Position
				cursor.AddChild(invalid)
			}
		}

		cursor.AddChild(node)
	} else {
		cursor = newLeaf(t)
		root = cursor
	}

	return root, cursor
}
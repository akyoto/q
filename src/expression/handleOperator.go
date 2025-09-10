package expression

// handleOperator adds a new operator to the tree.
func handleOperator(root *Expression, cursor *Expression, node *Expression) *Expression {
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

	return root
}
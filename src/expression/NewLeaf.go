package expression

import "git.urbach.dev/cli/q/src/token"

// NewLeaf creates a new leaf node.
func NewLeaf(t token.Token) *Expression {
	return &Expression{Token: t}
}
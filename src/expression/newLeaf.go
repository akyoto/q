package expression

import "git.urbach.dev/cli/q/src/token"

// newLeaf creates a new leaf node.
func newLeaf(t token.Token) *Expression {
	return &Expression{Token: t}
}
package expression

import (
	"git.urbach.dev/cli/q/src/token"
)

// coverGroup sets the token of `node` to cover the whole group.
func coverGroup(node *Expression, tokens token.List, groupPosition uint, t token.Token) {
	open := tokens[groupPosition-1]
	node.Token.Position = open.Position
	node.Token.Length = token.Length(t.End() - open.Position)
}
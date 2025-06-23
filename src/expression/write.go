package expression

import (
	"strings"

	"git.urbach.dev/cli/q/src/token"
)

// write generates a textual representation of the expression.
func (expr *Expression) write(builder *strings.Builder, source []byte) {
	if expr.IsLeaf() {
		builder.WriteString(expr.Token.String(source))
		return
	}

	builder.WriteByte('(')

	switch expr.Token.Kind {
	case token.Call:
		builder.WriteString(Operators[token.Call].Symbol)
	case token.Array:
		builder.WriteString(Operators[token.Array].Symbol)
	default:
		builder.WriteString(expr.Token.String(source))
	}

	for _, child := range expr.Children {
		builder.WriteByte(' ')
		child.write(builder, source)
	}

	builder.WriteByte(')')
}
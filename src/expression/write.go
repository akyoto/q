package expression

import (
	"strings"

	"git.urbach.dev/cli/q/src/token"
)

// write generates a textual representation of the expression.
func (expr *Expression) write(builder *strings.Builder, source []byte) {
	if expr.IsLeaf() {
		builder.WriteString(expr.Token.StringFrom(source))
		return
	}

	builder.WriteByte('(')

	switch expr.Token.Kind {
	case token.Call:
		builder.WriteString(token.Call.String())
	case token.Array:
		builder.WriteString(token.Array.String())
	case token.Struct:
		builder.WriteString(token.Struct.String())
	default:
		builder.WriteString(expr.Token.StringFrom(source))
	}

	for _, child := range expr.Children {
		builder.WriteByte(' ')
		child.write(builder, source)
	}

	builder.WriteByte(')')
}
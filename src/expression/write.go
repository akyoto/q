package expression

import (
	"strings"
)

// write generates a textual representation of the expression.
func (expr *Expression) write(builder *strings.Builder, source []byte) {
	if expr.IsLeaf() {
		builder.WriteString(expr.Token.StringFrom(source))
		return
	}

	builder.WriteByte('(')
	builder.WriteString(expr.Token.Kind.String())

	for _, child := range expr.Children {
		builder.WriteByte(' ')
		child.write(builder, source)
	}

	builder.WriteByte(')')
}
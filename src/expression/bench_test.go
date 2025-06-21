package expression_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
)

func BenchmarkExpression(b *testing.B) {
	src := []byte("(1+2-3*4)+(5*6-7+8)")
	tokens := token.Tokenize(src)

	for b.Loop() {
		expression.Parse(tokens)
	}
}
package expression_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

func TestFromTokens(t *testing.T) {
	tests := []struct {
		Name       string
		Expression string
		Result     string
	}{
		{"Identity", "1", "1"},
		{"Basic calculation", "1+2", "(1+2)"},
		{"Same operator", "1+2+3", "((1+2)+3)"},
		{"Same operator 2", "1+2+3+4", "(((1+2)+3)+4)"},
		{"Different operator", "1+2-3", "((1+2)-3)"},
		{"Different operator 2", "1+2-3+4", "(((1+2)-3)+4)"},
		{"Different operator 3", "1+2-3+4-5", "((((1+2)-3)+4)-5)"},
		{"Grouping identity", "(1)", "1"},
		{"Grouping identity 2", "((1))", "1"},
		{"Grouping identity 3", "(((1)))", "1"},
		{"Adding identity", "(1)+(2)", "(1+2)"},
		{"Adding identity 2", "(1)+(2)+(3)", "((1+2)+3)"},
		{"Adding identity 3", "(1)+(2)+(3)+(4)", "(((1+2)+3)+4)"},
		{"Grouping", "(1+2)", "(1+2)"},
		{"Grouping 2", "(1+2+3)", "((1+2)+3)"},
		{"Grouping 3", "((1)+(2)+(3))", "((1+2)+3)"},
		{"Grouping left", "(1+2)*3", "((1+2)*3)"},
		{"Grouping right", "1*(2+3)", "(1*(2+3))"},
		{"Grouping same operator", "1+(2+3)", "(1+(2+3))"},
		{"Grouping same operator 2", "1+(2+3)+(4+5)", "((1+(2+3))+(4+5))"},
		{"Two groups", "(1+2)*(3+4)", "((1+2)*(3+4))"},
		{"Two groups 2", "(1+2-3)*(3+4-5)", "(((1+2)-3)*((3+4)-5))"},
		{"Two groups 3", "(1+2)*(3+4-5)", "((1+2)*((3+4)-5))"},
		{"Operator priority", "1+2*3", "(1+(2*3))"},
		{"Operator priority 2", "1*2+3", "((1*2)+3)"},
		{"Operator priority 3", "1+2*3+4", "((1+(2*3))+4)"},
		{"Operator priority 4", "1+2*3*4", "((1+(2*3))*4)"},
		{"Operator priority 5", "1+2*3+4*5", "((1+(2*3))+(4*5))"},
		{"Complex", "(1+2-3*4)*(5+6-7*8)", "(((1+2)-(3*4))*((5+6)-(7*8)))"},
		{"Complex 2", "(1+2*3-4)*(5+6*7-8)", "(((1+(2*3))-4)*((5+(6*7))-8))"},
		{"Complex 3", "(1+2*3-4)*(5+6*7-8)+9-10*11", "(((((1+(2*3))-4)*((5+(6*7))-8))+9)-(10*11))"},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			src := []byte(test.Expression + "\n")

			tokens, processed := token.Tokenize(src, []token.Token{})
			assert.Equal(t, processed, len(src))

			expr, err := expression.FromTokens(tokens)
			assert.Nil(t, err)
			assert.NotNil(t, expr)
			assert.Equal(t, expr.String(), test.Result)
		})
	}
}

func BenchmarkExpression(b *testing.B) {
	src := []byte("(1+2-3*4)*(5+6-7*8)\n")
	tokens, _ := token.Tokenize(src, []token.Token{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = expression.FromTokens(tokens)
	}
}

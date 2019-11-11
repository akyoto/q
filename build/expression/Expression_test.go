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
		{"Same operator", "1+2+3", "(1+2+3)"},
		{"Different operator", "1+2-3", "((1+2)-3)"},
		{"Operator precedence", "1+2*3", "(1+(2*3))"},
		{"Grouping identity", "(1)", "1"},
		{"Adding identity 2 times", "(1)+(2)", "(1+2)"},
		{"Adding identity 3 times", "(1)+(2)+(3)", "(1+2+3)"},
		// {"Grouping left", "(1+2)*3", "((1+2)*3)"},
		{"Grouping right", "1*(2+3)", "(1*(2+3))"},
		// {"Grouping same operator", "1+(2+3)", "(1+2+3)"},
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

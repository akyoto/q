package expression_test

import (
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

func TestExpressionFromTokens(t *testing.T) {
	tests := []struct {
		Name       string
		Expression string
		Result     string
	}{
		{"Empty", "", ""},
		{"Identity", "1", "1"},
		{"Basic calculation", "1+2", "(1+2)"},
		{"Same operator", "1+2+3", "((1+2)+3)"},
		{"Same operator 2", "1+2+3+4", "(((1+2)+3)+4)"},
		{"Different operator", "1+2-3", "((1+2)-3)"},
		{"Different operator 2", "1+2-3+4", "(((1+2)-3)+4)"},
		{"Different operator 3", "1+2-3+4-5", "((((1+2)-3)+4)-5)"},
		{"Grouped identity", "(1)", "1"},
		{"Grouped identity 2", "((1))", "1"},
		{"Grouped identity 3", "(((1)))", "1"},
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
		{"Operator priority 4", "1+2*(3+4)+5", "((1+(2*(3+4)))+5)"},
		{"Operator priority 5", "1+2*3*4", "(1+((2*3)*4))"},
		{"Operator priority 6", "1+2*3+4*5", "((1+(2*3))+(4*5))"},
		{"Operator priority 7", "1+2*3*4*5*6", "(1+((((2*3)*4)*5)*6))"},
		{"Operator priority 8", "1*2*3+4*5*6", "(((1*2)*3)+((4*5)*6))"},
		{"Complex", "(1+2-3*4)*(5+6-7*8)", "(((1+2)-(3*4))*((5+6)-(7*8)))"},
		{"Complex 2", "(1+2*3-4)*(5+6*7-8)", "(((1+(2*3))-4)*((5+(6*7))-8))"},
		{"Complex 3", "(1+2*3-4)*(5+6*7-8)+9-10*11", "(((((1+(2*3))-4)*((5+(6*7))-8))+9)-(10*11))"},
		{"Function calls", "a()", "a()"},
		{"Function calls 2", "a(1)", "a(1)"},
		{"Function calls 3", "a(1,2)", "a(1,2)"},
		{"Function calls 4", "a(1,2,3)", "a(1,2,3)"},
		{"Function calls 5", "a(1,2+2,3)", "a(1,(2+2),3)"},
		{"Function calls 6", "a(1,2+2,3+3)", "a(1,(2+2),(3+3))"},
		{"Function calls 7", "a(1+1,2,3)", "a((1+1),2,3)"},
		{"Function calls 8", "a(1+1,2+2,3+3)", "a((1+1),(2+2),(3+3))"},
		{"Function calls 9", "a(b())", "a(b())"},
		{"Function calls 10", "a(b(),c())", "a(b(),c())"},
		{"Function calls 11", "a(b(),c(),d())", "a(b(),c(),d())"},
		{"Function calls 12", "a(b(1),c(2),d(3))", "a(b(1),c(2),d(3))"},
		{"Function calls 13", "a(b(1)+1)", "a((b(1)+1))"},
		{"Function calls 14", "a(b(1)+1,c(2),d(3))", "a((b(1)+1),c(2),d(3))"},
		{"Function calls 15", "a(b(1)*c(2))", "a((b(1)*c(2)))"},
		{"Function calls 16", "a(b(1)*c(2),d(3)+e(4),f(5)/f(6))", "a((b(1)*c(2)),(d(3)+e(4)),(f(5)/f(6)))"},
		{"Function calls 17", "a((b(1,2)+c(3,4))*d(5,6))", "a(((b(1,2)+c(3,4))*d(5,6)))"},
		{"Function calls 18", "a((b(1,2)+c(3,4))*d(5,6),e())", "a(((b(1,2)+c(3,4))*d(5,6)),e())"},
		{"Function calls 19", "a((b(1,2)+c(3,4))*d(5,6),e(7+8,9-10*11,12))", "a(((b(1,2)+c(3,4))*d(5,6)),e((7+8),(9-(10*11)),12))"},
		{"Function calls 20", "a((b(1,2,bb())+c(3,4,cc(0)))*d(5,6,dd(0)),e(7+8,9-10*11,12,ee(0)))", "a(((b(1,2,bb())+c(3,4,cc(0)))*d(5,6,dd(0))),e((7+8),(9-(10*11)),12,ee(0)))"},
		{"Function calls 21", "a(1-2*3)", "a((1-(2*3)))"},
		{"Function calls 22", "1+2*a()+4", "((1+(2*a()))+4)"},
		{"Function calls 23", "sum(a,b)*2+15*4", "((sum(a,b)*2)+(15*4))"},
		{"Package function calls", "math.sum(a,b)", "(math.sum(a,b))"},
		{"Package function calls 2", "generic.math.sum(a,b)", "((generic.math).sum(a,b))"},
	}

	for _, test := range tests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			src := []byte(test.Expression + "\n")

			tokens, processed := token.Tokenize(src, []token.Token{})
			assert.Equal(t, processed, uint16(len(src)))
			tokens = tokens[:len(tokens)-1]

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
		expr, _ := expression.FromTokens(tokens)
		expr.Close()
	}
}

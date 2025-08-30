package expression_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		Name       string
		Expression string
		Result     string
	}{
		{"Identity", "1", "1"},
		{"Basic calculation", "1+2", "(+ 1 2)"},

		{"Same operator", "1+2+3", "(+ (+ 1 2) 3)"},
		{"Same operator 2", "1+2+3+4", "(+ (+ (+ 1 2) 3) 4)"},

		{"Different operator", "1+2-3", "(- (+ 1 2) 3)"},
		{"Different operator 2", "1+2-3+4", "(+ (- (+ 1 2) 3) 4)"},
		{"Different operator 3", "1+2-3+4-5", "(- (+ (- (+ 1 2) 3) 4) 5)"},

		{"Grouped identity", "(1)", "1"},
		{"Grouped identity 2", "((1))", "1"},
		{"Grouped identity 3", "(((1)))", "1"},

		{"Adding identity", "(1)+(2)", "(+ 1 2)"},
		{"Adding identity 2", "(1)+(2)+(3)", "(+ (+ 1 2) 3)"},
		{"Adding identity 3", "(1)+(2)+(3)+(4)", "(+ (+ (+ 1 2) 3) 4)"},

		{"Grouping", "(1+2)", "(+ 1 2)"},
		{"Grouping 2", "(1+2+3)", "(+ (+ 1 2) 3)"},
		{"Grouping 3", "((1)+(2)+(3))", "(+ (+ 1 2) 3)"},
		{"Grouping left", "(1+2)*3", "(* (+ 1 2) 3)"},
		{"Grouping right", "1*(2+3)", "(* 1 (+ 2 3))"},
		{"Grouping same operator", "1+(2+3)", "(+ 1 (+ 2 3))"},
		{"Grouping same operator 2", "1+(2+3)+(4+5)", "(+ (+ 1 (+ 2 3)) (+ 4 5))"},

		{"Two groups", "(1+2)*(3+4)", "(* (+ 1 2) (+ 3 4))"},
		{"Two groups 2", "(1+2-3)*(3+4-5)", "(* (- (+ 1 2) 3) (- (+ 3 4) 5))"},
		{"Two groups 3", "(1+2)*(3+4-5)", "(* (+ 1 2) (- (+ 3 4) 5))"},

		{"Operator priority", "1+2*3", "(+ 1 (* 2 3))"},
		{"Operator priority 2", "1*2+3", "(+ (* 1 2) 3)"},
		{"Operator priority 3", "1+2*3+4", "(+ (+ 1 (* 2 3)) 4)"},
		{"Operator priority 4", "1+2*(3+4)+5", "(+ (+ 1 (* 2 (+ 3 4))) 5)"},
		{"Operator priority 5", "1+2*3*4", "(+ 1 (* (* 2 3) 4))"},
		{"Operator priority 6", "1+2*3+4*5", "(+ (+ 1 (* 2 3)) (* 4 5))"},
		{"Operator priority 7", "1+2*3*4*5*6", "(+ 1 (* (* (* (* 2 3) 4) 5) 6))"},
		{"Operator priority 8", "1*2*3+4*5*6", "(+ (* (* 1 2) 3) (* (* 4 5) 6))"},

		{"Complex", "(1+2-3*4)*(5+6-7*8)", "(* (- (+ 1 2) (* 3 4)) (- (+ 5 6) (* 7 8)))"},
		{"Complex 2", "(1+2*3-4)*(5+6*7-8)", "(* (- (+ 1 (* 2 3)) 4) (- (+ 5 (* 6 7)) 8))"},
		{"Complex 3", "(1+2*3-4)*(5+6*7-8)+9-10*11", "(- (+ (* (- (+ 1 (* 2 3)) 4) (- (+ 5 (* 6 7)) 8)) 9) (* 10 11))"},

		{"Unary not", "!", "!"},
		{"Unary not 2", "!a", "(! a)"},
		{"Unary not 3", "!(!a)", "(! (! a))"},
		{"Unary not 4", "!(a||b)", "(! (|| a b))"},
		{"Unary not 5", "a || !b", "(|| a (! b))"},

		{"Unary minus", "-", "-"},
		{"Unary minus 2", "-a", "(- a)"},
		{"Unary minus 3", "-(-a)", "(- (- a))"},
		{"Unary minus 4", "-a+b", "(+ (- a) b)"},
		{"Unary minus 5", "-(a+b)", "(- (+ a b))"},
		{"Unary minus 6", "a + -b", "(+ a (- b))"},
		{"Unary minus 7", "-a + -b", "(+ (- a) (- b))"},

		{"Assign bitwise operation", "a|=b", "(|= a b)"},
		{"Assign bitwise operation 2", "a|=b<<c", "(|= a (<< b c))"},

		{"Function calls", "a()", "(λ a)"},
		{"Function calls 2", "a(1)", "(λ a 1)"},
		{"Function calls 3", "a(1)+1", "(+ (λ a 1) 1)"},
		{"Function calls 4", "1+a(1)", "(+ 1 (λ a 1))"},
		{"Function calls 5", "a(1,2)", "(λ a 1 2)"},
		{"Function calls 6", "a(1,2,3)", "(λ a 1 2 3)"},
		{"Function calls 7", "a(1,2+2,3)", "(λ a 1 (+ 2 2) 3)"},
		{"Function calls 8", "a(1,2+2,3+3)", "(λ a 1 (+ 2 2) (+ 3 3))"},
		{"Function calls 9", "a(1+1,2,3)", "(λ a (+ 1 1) 2 3)"},
		{"Function calls 10", "a(1+1,2+2,3+3)", "(λ a (+ 1 1) (+ 2 2) (+ 3 3))"},
		{"Function calls 11", "a(b())", "(λ a (λ b))"},
		{"Function calls 12", "a(b(),c())", "(λ a (λ b) (λ c))"},
		{"Function calls 13", "a(b(),c(),d())", "(λ a (λ b) (λ c) (λ d))"},
		{"Function calls 14", "a(b(1))", "(λ a (λ b 1))"},
		{"Function calls 15", "a(b(1),c(2),d(3))", "(λ a (λ b 1) (λ c 2) (λ d 3))"},
		{"Function calls 16", "a(b(1)+1)", "(λ a (+ (λ b 1) 1))"},
		{"Function calls 17", "a(b(1)+1,c(2),d(3))", "(λ a (+ (λ b 1) 1) (λ c 2) (λ d 3))"},
		{"Function calls 18", "a(b(1)*c(2))", "(λ a (* (λ b 1) (λ c 2)))"},
		{"Function calls 19", "a(b(1)*c(2),d(3)+e(4),f(5)/f(6))", "(λ a (* (λ b 1) (λ c 2)) (+ (λ d 3) (λ e 4)) (/ (λ f 5) (λ f 6)))"},
		{"Function calls 20", "a(b(1,2)+c(3,4)*d(5,6))", "(λ a (+ (λ b 1 2) (* (λ c 3 4) (λ d 5 6))))"},
		{"Function calls 21", "a((b(1,2)+c(3,4))*d(5,6))", "(λ a (* (+ (λ b 1 2) (λ c 3 4)) (λ d 5 6)))"},
		{"Function calls 22", "a((b(1,2)+c(3,4))*d(5,6),e())", "(λ a (* (+ (λ b 1 2) (λ c 3 4)) (λ d 5 6)) (λ e))"},
		{"Function calls 23", "a((b(1,2)+c(3,4))*d(5,6),e(7+8,9-10*11,12))", "(λ a (* (+ (λ b 1 2) (λ c 3 4)) (λ d 5 6)) (λ e (+ 7 8) (- 9 (* 10 11)) 12))"},
		{"Function calls 24", "a((b(1,2,bb())+c(3,4,cc(0)))*d(5,6,dd(0)),e(7+8,9-10*11,12,ee(0)))", "(λ a (* (+ (λ b 1 2 (λ bb)) (λ c 3 4 (λ cc 0))) (λ d 5 6 (λ dd 0))) (λ e (+ 7 8) (- 9 (* 10 11)) 12 (λ ee 0)))"},
		{"Function calls 25", "a(1-2*3)", "(λ a (- 1 (* 2 3)))"},
		{"Function calls 26", "1+2*a()+4", "(+ (+ 1 (* 2 (λ a))) 4)"},
		{"Function calls 27", "a(b,c)*2+15*4", "(+ (* (λ a b c) 2) (* 15 4))"},
		{"Function calls 28", "a(,)", "(λ a  )"},

		{"Package function calls", "a.b(c)", "(λ (. a b) c)"},
		{"Package function calls 2", "a.b(c,d)", "(λ (. a b) c d)"},
		{"Package function calls 3", "a.b.c(d,e)", "(λ (. (. a b) c) d e)"},

		{"Array access", "a[0]", "(@ a 0)"},
		{"Array access 2", "a[b+c]", "(@ a (+ b c))"},
		{"Array access 3", "a.b[c]", "(@ (. a b) c)"},
		{"Array access 4", "a.b[c+d]", "(@ (. a b) (+ c d))"},
		{"Array access 5", "a()[b]", "(@ (λ a) b)"},
		{"Array access 6", "a.b()[c]", "(@ (λ (. a b)) c)"},
		{"Array access 7", "a.b(c)[d]", "(@ (λ (. a b) c) d)"},
		{"Array access 8", "a.b(c)[d][e]", "(@ (@ (λ (. a b) c) d) e)"},
		{"Array access 9", "a[0](1)[2](3)", "(λ (@ (λ (@ a 0) 1) 2) 3)"},

		{"Dereferencing", "[a]", "(@ a)"},
		{"Dereferencing 2", "[a+b]", "(@ (+ a b))"},
		{"Dereferencing 3", "[a+b]=c", "(= (@ (+ a b)) c)"},
		{"Dereferencing 4", "[a+b]=c+d", "(= (@ (+ a b)) (+ c d))"},

		{"Structs", "a{}", "($ a)"},
		{"Structs 2", "a{b:c}", "($ a (: b c))"},
		{"Structs 3", "a{b:c,d:e}", "($ a (: b c) (: d e))"},
		{"Structs 4", "a{}.f", "(. ($ a) f)"},
		{"Structs 5", "a{b:c}.f", "(. ($ a (: b c)) f)"},
		{"Structs 6", "a{}[0]", "(@ ($ a) 0)"},
		{"Structs 7", "a{b:c}[0]", "(@ ($ a (: b c)) 0)"},
		{"Structs 8", "a{}(f)", "(λ ($ a) f)"},
		{"Structs 9", "a{b:c}(f)", "(λ ($ a (: b c)) f)"},

		{"Slices", "a[..]", "(@ a ..)"},
		{"Slices 2", "a[b..]", "(@ a (.. b))"},
		{"Slices 3", "a[..c]", "(@ a (..  c))"},
		{"Slices 4", "a[b..c]", "(@ a (.. b c))"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			src := []byte(test.Expression)
			tokens := token.Tokenize(src)
			expr := expression.Parse(tokens)
			defer expr.Reset()
			assert.NotNil(t, expr)
			assert.Equal(t, expr.String(src), test.Result)
		})
	}
}
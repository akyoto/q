package tests_test

import (
	"testing"
)

var tests = []testRun{
	{"empty", "", "", 0},
	{"assert", "", "", 1},
	{"binary", "", "", 0},
	{"octal", "", "", 0},
	{"hexadecimal", "", "", 0},
	{"rune", "", "", 0},
	{"variables", "", "", 0},
	{"reuse", "", "", 0},
	{"return", "", "", 0},
	{"add", "", "", 0},
	{"sub", "", "", 0},
	{"mul", "", "", 0},
	{"div", "", "", 0},
	{"sum", "", "", 0},
	{"square-sum", "", "", 0},
	{"math", "", "", 0},
	{"bitwise-and", "", "", 0},
	{"bitwise-or", "", "", 0},
	{"bitwise-xor", "", "", 0},
	{"modulo", "", "", 0},
	{"negative", "", "", 0},
	{"negation", "", "", 0},
	{"param", "", "", 0},
	{"param-multi", "", "", 0},
	{"param-order", "", "", 0},
	{"branch", "", "", 0},
	{"branch-and", "", "", 0},
	{"branch-or", "", "", 0},
	{"branch-both", "", "", 0},
	{"jump-near", "", "", 0},
	{"hello", "", "Hello\nHello\nHello\n", 0},
	{"script", "", "Hello\n", 0},
}

func TestTests(t *testing.T) {
	for _, test := range tests {
		test.Run(t, test.Name+".q")
	}
}
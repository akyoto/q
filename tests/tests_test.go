package tests_test

import (
	"testing"
)

var tests = []run{
	{"empty", "", "", 0},
	{"assert", "", "", 1},
	{"binary", "", "", 0},
	{"octal", "", "", 0},
	{"hexadecimal", "", "", 0},
	{"rune", "", "", 0},
	{"variables", "", "", 0},
	{"reassign", "", "", 0},
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
	{"shift", "", "", 0},
	{"modulo", "", "", 0},
	{"negative", "", "", 0},
	{"negation", "", "", 0},
	{"assign", "", "", 0},
	{"param", "", "", 0},
	{"param-multi", "", "", 0},
	{"param-order", "", "", 0},
	{"branch", "", "", 0},
	{"branch-and", "", "", 0},
	{"branch-or", "", "", 0},
	{"branch-both", "", "", 0},
	{"jump-near", "", "", 0},
	{"switch", "", "", 0},
	{"phi", "", "", 0},
	{"phi-simple", "", "", 0},
	{"phi-advanced", "", "", 0},
	{"loop", "", "..........", 0},
	{"loop-stacked", "", "..........", 0},
	{"loop-count", "", "", 0},
	{"loop-write", "", "..........", 0},
	{"loop-optimize-single-iterator", "", "", 0},
	{"loop-limit", "", "", 0},
	{"loop-limit-dynamic", "", "", 0},
	{"loop-keepalive", "", "", 0},
	{"memory", "", "Hello\n", 0},
	{"out-of-memory", "", "", 1},
	{"index-static", "", "", 0},
	{"index-dynamic", "", "", 0},
	{"struct", "", "", 0},
	{"struct-init", "", "", 0},
	{"custom-string", "", "012345", 0},
	{"return-2", "", "", 0},
	{"return-3", "", "", 0},
	{"return-4", "", "", 0},
	{"return-string", "", "Hello\n", 0},
	{"ignore-unused-field", "", "", 0},
	{"hello", "", "Hello\nHello\nHello\n", 0},
	{"escape", "", "a\tb\nc\td\n", 0},
	{"script", "", "Hello\n", 0},
}

func TestTests(t *testing.T) {
	for _, test := range tests {
		test.Run(t, test.Name+".q")
	}
}
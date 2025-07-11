package tests_test

import (
	"testing"
)

var tests = []testRun{
	{"script", "", "Hello\n", 0},
	{"sum-10", "", "", 10},
	{"sum-36", "", "", 36},
	{"swap", "", "", 3},
	{"value-reuse", "", "Hello\nHello\nHello\n", 0},
}

func TestTests(t *testing.T) {
	for _, test := range tests {
		test.Run(t, test.Name+".q")
	}
}
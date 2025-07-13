package tests_test

import (
	"testing"
)

var tests = []testRun{
	{"empty", "", "", 0},
	{"sum-10", "", "", 10},
	{"sum-36", "", "", 36},
	{"hello-3", "", "Hello\nHello\nHello\n", 0},
	{"param-swap", "", "", 3},
	{"script", "", "Hello\n", 0},
}

func TestTests(t *testing.T) {
	for _, test := range tests {
		test.Run(t, test.Name+".q")
	}
}
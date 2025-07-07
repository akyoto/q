package tests_test

import (
	"testing"
)

var tests = []testRun{
	{"script", "", "Hello\n", 0},
	{"sum", "", "", 10},
}

func TestTests(t *testing.T) {
	for _, test := range tests {
		test.Run(t, test.Name+".q")
	}
}
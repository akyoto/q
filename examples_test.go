package main_test

import (
	"testing"
)

// examples is a list of examples with their expected output and exit code.
// This list is used in multiple tests.
var examples = []struct {
	Name             string
	ExpectedOutput   string
	ExpectedExitCode int
}{
	{"hello", "Hello\n", 0},
	{"contracts", "f: expect [n < 10]\n", 1},
	{"fibonacci", "", 89},
	{"files", "", 0},
	{"functions", "123456789\n123456789\n123456789\n123456789\n", 0},
	{"loops", "Hello\nHello\nHello\n\nH\nHe\nHel\nHell\nHello\n", 0},
	{"memory", "ABCD\n", 0},
	{"struct", "", 100},
}

func TestExamples(t *testing.T) {
	for _, example := range examples {
		example := example

		t.Run(example.Name, func(t *testing.T) {
			Run(t, "./examples/"+example.Name, example.ExpectedOutput, example.ExpectedExitCode)
		})
	}
}

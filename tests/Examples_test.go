package build_test

import (
	"testing"
)

func TestExamples(t *testing.T) {
	examples := []struct {
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
	}

	for _, example := range examples {
		example := example

		t.Run(example.Name, func(t *testing.T) {
			Run(t, "../examples/"+example.Name, example.ExpectedOutput, example.ExpectedExitCode)
		})
	}
}

package build_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/cli"
)

func TestCLI(t *testing.T) {
	defer os.Remove("../examples/hello/hello")

	os.Args = []string{"q", "build", "../examples/hello"}
	cli.Main()

	stat, err := os.Stat("../examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}

func TestCLIHelp(t *testing.T) {
	cli.Help()
}

func TestExamples(t *testing.T) {
	examples := []struct {
		Name           string
		ExpectedOutput string
	}{
		{"hello", "Hello\n"},
		{"functions", "Function 1\nFunction 2\nFunction 3\n"},
		{"syscalls", "Hello Syscalls"},
	}

	for _, example := range examples {
		directory := "../examples/" + example.Name
		output := example.ExpectedOutput

		t.Run(example.Name, func(t *testing.T) {
			assertOutput(t, directory, output)
		})
	}
}

func TestBuildErrors(t *testing.T) {
	tests := []struct {
		File          string
		ExpectedError string
	}{
		{"testdata/missing-opening-bracket.q", "Missing opening bracket"},
		{"testdata/missing-closing-bracket.q", "Missing closing bracket"},
		{"testdata/unknown-function.q", "Unknown function"},
		{"testdata/unknown-function-suggestion.q", "Unknown function 'prin', did you mean 'print'?"},
	}

	for _, test := range tests {
		compiler, err := syntaxChecker(test.File)
		assert.Nil(t, err)
		err = compiler.Run()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), test.ExpectedError)
	}
}

func TestTokenizerErrors(t *testing.T) {
	tests := []struct {
		File          string
		ExpectedError string
	}{
		{"testdata/unknown-expression.q", "Unknown expression"},
	}

	for _, test := range tests {
		_, err := syntaxChecker(test.File)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), test.ExpectedError)
	}
}

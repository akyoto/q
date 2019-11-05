package build_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build"
)

func TestCLI(t *testing.T) {
	defer os.Remove("../examples/hello/hello")

	os.Args = []string{"q", "build", "../examples/hello"}
	build.Main()

	stat, err := os.Stat("../examples/hello/hello")
	assert.Nil(t, err)
	assert.True(t, stat.Size() > 0)
}

func TestCLIHelp(t *testing.T) {
	build.Help()
}

func TestExamplesHelloWorld(t *testing.T) {
	assertOutput(t, "../examples/hello", "Hello\n")
}

func TestExamplesFunctions(t *testing.T) {
	assertOutput(t, "../examples/functions", "Hello functions\n")
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
		file := build.NewFile(test.File)
		defer file.Close()
		err := file.Compile()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), test.ExpectedError)
	}
}

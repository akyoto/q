package compiler_test

import (
	"os"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/compiler"
)

func TestCLI(t *testing.T) {
	defer os.Remove("hello")
	os.Args = []string{"./q", "testdata/hello.q"}
	compiler.Main()
}

func TestCLIHelp(t *testing.T) {
	compiler.Help()
}

func TestCompiler(t *testing.T) {
	defer os.Remove("a.out")
	compiler := compiler.New()
	err := compiler.Compile("testdata/compiler-bench-1k.q", "a.out")
	assert.Nil(t, err)
}

func TestCompilerErrors(t *testing.T) {
	tests := []struct {
		File          string
		ExpectedError string
	}{
		{"testdata/syntax-errors/missing-opening-bracket.q", "Missing opening bracket"},
		{"testdata/syntax-errors/missing-closing-bracket.q", "Missing closing bracket"},
		{"testdata/syntax-errors/unknown-function.q", "Unknown function"},
		{"testdata/syntax-errors/unknown-function-suggestion.q", "Unknown function 'prin', did you mean 'print'?"},
	}

	for _, test := range tests {
		compiler := compiler.New()
		defer compiler.Close()
		err := compiler.Compile(test.File, "a.out")
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), test.ExpectedError)
	}
}

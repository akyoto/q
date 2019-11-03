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

func TestSyntaxErrorMissingOpeningBracket(t *testing.T) {
	compiler := compiler.New()
	err := compiler.Compile("testdata/syntax-errors/missing-opening-bracket.q", "a.out")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Missing opening bracket")
}

func TestSyntaxErrorMissingClosingBracket(t *testing.T) {
	compiler := compiler.New()
	err := compiler.Compile("testdata/syntax-errors/missing-closing-bracket.q", "a.out")
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Missing closing bracket")
}

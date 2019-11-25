package build_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/log"
)

func TestMain(m *testing.M) {
	log.Info.SetOutput(ioutil.Discard)
	log.Error.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestExamples(t *testing.T) {
	examples := []struct {
		Name             string
		ExpectedOutput   string
		ExpectedExitCode int
	}{
		{"hello", "Hello\n", 0},
		{"procedures", "Procedure 1\nProcedure 2\nProcedure 3\n", 0},
		{"syscalls", "Hello Syscalls", 0},
		{"variables", "", 100},
		{"loops", "Hello\nHello\nHello\n", 0},
	}

	for _, example := range examples {
		example := example

		t.Run(example.Name, func(t *testing.T) {
			assertOutput(t, "../examples/"+example.Name, example.ExpectedOutput, example.ExpectedExitCode)
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
		{"testdata/unused-variable.q", "Variable 'a' has never been used"},
		{"testdata/unknown-function.q", "Unknown function"},
		{"testdata/unknown-function-suggestion.q", "Unknown function 'prin', did you mean 'print'?"},
		{"testdata/unknown-expression.q", "Unknown expression"},
	}

	for _, test := range tests {
		test := test
		name := strings.TrimPrefix(test.File, "testdata/")
		name = strings.TrimSuffix(name, ".q")

		t.Run(name, func(t *testing.T) {
			err := syntaxCheck(test.File)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError)
		})
	}
}

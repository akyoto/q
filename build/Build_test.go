package build_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/errors"
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
		{"functions", "123456789\n", 0},
		{"syscalls", "123456789\n", 0},
		{"variables", "", 100},
		{"loops", "Hello\nHello\nHello\n", 0},
		{"fibonacci", "", 89},
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
		ExpectedError error
	}{
		{"testdata/missing-opening-bracket.q", &errors.MissingCharacter{Character: "("}},
		{"testdata/missing-closing-bracket.q", &errors.MissingCharacter{Character: ")"}},
		{"testdata/unused-variable.q", &errors.UnusedVariable{VariableName: "a"}},
		{"testdata/unknown-function.q", &errors.UnknownFunction{FunctionName: "z"}},
		{"testdata/unknown-function-suggestion.q", &errors.UnknownFunction{FunctionName: "prin", CorrectName: "print"}},
		{"testdata/unknown-expression.q", &errors.UnknownExpression{Expression: "\")"}},
		{"testdata/for-missing-upper-limit.q", errors.MissingRangeLimit},
		{"testdata/for-missing-range.q", errors.MissingRange},
		{"testdata/for-missing-start-value.q", errors.MissingRangeStart},
	}

	for _, test := range tests {
		test := test
		name := strings.TrimPrefix(test.File, "testdata/")
		name = strings.TrimSuffix(name, ".q")

		t.Run(name, func(t *testing.T) {
			err := syntaxCheck(test.File)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}

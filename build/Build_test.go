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
		{"contracts", "f: require [n < 10]\n", 1},
		{"fibonacci", "", 89},
		{"files", "", 0},
		{"functions", "123456789\n123456789\n123456789\n123456789\n", 0},
		{"loops", "Hello\nHello\nHello\n\nH\nHe\nHel\nHell\nHello\n", 0},
		{"memory", "", 0},
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
		{"testdata/for-missing-upper-limit.q", errors.MissingRangeLimit},
		{"testdata/for-missing-range.q", errors.MissingRange},
		{"testdata/for-missing-start-value.q", errors.MissingRangeStart},
		{"testdata/ineffective-assignment.q", &errors.IneffectiveAssignment{Name: "a"}},
		{"testdata/missing-opening-bracket.q", &errors.MissingCharacter{Character: "("}},
		{"testdata/missing-closing-bracket.q", &errors.MissingCharacter{Character: ")"}},
		{"testdata/missing-return-type.q", errors.MissingReturnType},
		{"testdata/missing-type.q", &errors.MissingType{Of: "length"}},
		{"testdata/return-without-type.q", errors.ReturnWithoutFunctionType},
		{"testdata/unnecessary-newlines.q", errors.UnnecessaryNewlines},
		{"testdata/unused-variable.q", &errors.UnusedVariable{Name: "a"}},
		{"testdata/unused-mutable.q", &errors.UnmodifiedMutable{Name: "a"}},
		{"testdata/unknown-function.q", &errors.UnknownFunction{Name: "z"}},
		{"testdata/unknown-function-suggestion.q", &errors.UnknownFunction{Name: "prin", CorrectName: "print"}},
		{"testdata/unknown-expression.q", &errors.UnknownExpression{Expression: "\")"}},
		{"testdata/unused-parameter.q", &errors.UnusedVariable{Name: "b"}},
	}

	for _, test := range tests {
		test := test
		name := strings.TrimPrefix(test.File, "testdata/")
		name = strings.TrimSuffix(name, ".q")

		t.Run(name, func(t *testing.T) {
			err := check(test.File)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}

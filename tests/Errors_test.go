package build_test

import (
	"strings"
	"testing"

	"github.com/akyoto/assert"
	"github.com/akyoto/q/build/errors"
)

func TestErrors(t *testing.T) {
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
		{"testdata/package-doesnt-exist.q", &errors.PackageDoesntExist{ImportPath: "non.existing.package"}},
		{"testdata/return-without-type.q", errors.ReturnWithoutFunctionType},
		{"testdata/unnecessary-newlines.q", errors.UnnecessaryNewlines},
		{"testdata/unused-variable.q", &errors.UnusedVariable{Name: "a"}},
		{"testdata/unused-mutable.q", &errors.UnmodifiedMutable{Name: "a"}},
		{"testdata/unknown-function.q", &errors.UnknownFunction{Name: "z"}},
		{"testdata/unknown-function-suggestion.q", &errors.UnknownFunction{Name: "prin", CorrectName: "print"}},
		{"testdata/unknown-expression.q", &errors.UnknownExpression{Expression: "\")"}},
		{"testdata/unknown-variable.q", &errors.UnknownVariable{Name: "a"}},
		{"testdata/unused-parameter.q", &errors.UnusedVariable{Name: "b"}},
		{"testdata/variable-already-exists.q", &errors.VariableAlreadyExists{Name: "a"}},
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

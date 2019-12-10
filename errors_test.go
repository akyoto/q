package main_test

import (
	"path/filepath"
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
		{"for-missing-upper-limit.q", errors.MissingRangeLimit},
		{"for-missing-range.q", errors.MissingRange},
		{"for-missing-start-value.q", errors.MissingRangeStart},
		{"immutable-variable.q", &errors.ImmutableVariable{Name: "a"}},
		{"import-already-exists.q", &errors.ImportNameAlreadyExists{Name: "sys", ImportPath: "sys"}},
		{"ineffective-assignment.q", &errors.IneffectiveAssignment{Name: "a"}},
		{"missing-opening-bracket.q", &errors.MissingCharacter{Character: "("}},
		{"missing-closing-bracket.q", &errors.MissingCharacter{Character: ")"}},
		{"missing-return-type.q", errors.MissingReturnType},
		{"missing-return-value.q", &errors.MissingReturnValue{ReturnType: "Int64"}},
		{"missing-type.q", &errors.MissingType{Of: "length"}},
		{"package-doesnt-exist.q", &errors.PackageDoesntExist{ImportPath: "non.existing.package"}},
		{"parameter-count.q", &errors.ParameterCount{FunctionName: "sum", CountGiven: 1, CountRequired: 2}},
		{"return-without-type.q", errors.ReturnWithoutFunctionType},
		{"unnecessary-newlines.q", errors.UnnecessaryNewlines},
		{"unused-variable.q", &errors.UnusedVariable{Name: "a"}},
		{"unused-mutable.q", &errors.UnmodifiedMutable{Name: "a"}},
		{"unknown-function.q", &errors.UnknownFunction{Name: "z"}},
		{"unknown-function-suggestion.q", &errors.UnknownFunction{Name: "prin", CorrectName: "print"}},
		{"unknown-expression.q", &errors.UnknownExpression{Expression: "\")"}},
		{"unknown-variable.q", &errors.UnknownVariable{Name: "a"}},
		{"unknown-variable-suggestion.q", &errors.UnknownVariable{Name: "lengt", CorrectName: "length"}},
		{"unknown-package.q", &errors.UnknownPackage{Name: "sy", CorrectName: "sys"}},
		{"unused-parameter.q", &errors.UnusedVariable{Name: "b"}},
		{"variable-already-exists.q", &errors.VariableAlreadyExists{Name: "a"}},
	}

	for _, test := range tests {
		test := test
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			err := Check(filepath.Join("build", "errors", "testdata", test.File))
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}

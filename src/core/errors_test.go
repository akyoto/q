package core_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"InvalidExpression.q", core.InvalidExpression},
	{"InvalidExpression2.q", core.InvalidExpression},
	{"InvalidExpression3.q", core.InvalidExpression},
	{"InvalidExpression4.q", core.InvalidExpression},
	{"InvalidLoopHeader.q", core.InvalidLoopHeader},
	{"MissingOperand.q", core.MissingOperand},
	{"ParameterCountMismatch.q", &core.ParameterCountMismatch{Function: "main.f", Count: 0, ExpectedCount: 1}},
	{"ParameterCountMismatch2.q", &core.ParameterCountMismatch{Function: "main.f", Count: 2, ExpectedCount: 1}},
	{"ReturnCountMismatch.q", &core.ReturnCountMismatch{Count: 1, ExpectedCount: 0}},
	{"ReturnCountMismatch2.q", &core.ReturnCountMismatch{Count: 1, ExpectedCount: 2}},
	{"TypeMismatch.q", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "x", IsReturn: false}},
	{"TypeMismatch2.q", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "y", IsReturn: true}},
	{"UnknownIdentifier.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier2.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier3.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier4.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier5.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier6.q", &core.UnknownIdentifier{Name: "os"}},
	{"UnknownIdentifier7.q", &core.UnknownIdentifier{Name: "os.unknown"}},
	{"UnknownIdentifier8.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier9.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownStructField.q", &core.UnknownStructField{StructName: "string", FieldName: "unknown"}},
	{"UnusedValue.q", &core.UnusedValue{Value: "42"}},
	{"UnusedValue2.q", &core.UnusedValue{Value: "2 + 3"}},
	{"UnusedValue3.q", &core.UnusedValue{Value: "\"not used\""}},
	{"UnusedValue4.q", &core.UnusedValue{Value: "1"}},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			b := config.New(filepath.Join("testdata", test.File))
			_, err := compiler.Compile(b)
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), test.ExpectedError.Error())
		})
	}
}
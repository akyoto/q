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
	{"InvalidExpression5.q", core.InvalidExpression},
	{"ParameterCountMismatch.q", &core.ParameterCountMismatch{Function: "main.f", Count: 0, ExpectedCount: 1}},
	{"ParameterCountMismatch2.q", &core.ParameterCountMismatch{Function: "main.f", Count: 2, ExpectedCount: 1}},
	{"TypeMismatch.q", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "x", IsReturn: false}},
	{"UnknownIdentifier.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier2.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier3.q", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier4.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier5.q", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier6.q", &core.UnknownIdentifier{Name: "os"}},
	{"UnknownIdentifier7.q", &core.UnknownIdentifier{Name: "os.unknown"}},
	{"UnusedValue.q", &core.UnusedValue{Value: "2 + 3"}},
	{"UnusedValue2.q", &core.UnusedValue{Value: "\"not used\""}},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			b := config.New(filepath.Join("testdata", test.File))
			_, err := compiler.Compile(b)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}
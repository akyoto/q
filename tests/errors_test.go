package tests_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	Name          string
	ExpectedError error
}{
	{"ParameterCountMismatch", &core.ParameterCountMismatch{Function: "main.f", Count: 0, ExpectedCount: 1}},
	{"ParameterCountMismatch2", &core.ParameterCountMismatch{Function: "main.f", Count: 2, ExpectedCount: 1}},
	{"TypeMismatch", &core.TypeMismatch{Encountered: "string", Expected: "int64", ParameterName: "x", IsReturn: false}},
	{"UnknownIdentifier", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier2", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier3", &core.UnknownIdentifier{Name: "x"}},
	{"UnknownIdentifier4", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier5", &core.UnknownIdentifier{Name: "unknown"}},
	{"UnknownIdentifier6", &core.UnknownIdentifier{Name: "os"}},
	{"UnknownIdentifier7", &core.UnknownIdentifier{Name: "os.unknown"}},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		t.Run(test.Name, func(t *testing.T) {
			build := config.New(filepath.Join("errors", test.Name+".q"))
			_, err := compiler.Compile(build)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}
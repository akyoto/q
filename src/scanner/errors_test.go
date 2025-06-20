package scanner_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"ExpectedFunctionDefinition.q", errors.ExpectedFunctionDefinition},
	{"ExpectedPackageName.q", errors.ExpectedPackageName},
	{"InvalidCharacter.q", &errors.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter2.q", &errors.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter3.q", &errors.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter4.q", &errors.InvalidCharacter{Character: "+++"}},
	{"InvalidFunctionDefinition.q", errors.InvalidFunctionDefinition},
	{"InvalidTopLevel.q", &errors.InvalidTopLevel{Instruction: "123"}},
	{"InvalidTopLevel2.q", &errors.InvalidTopLevel{Instruction: "\"Hello\""}},
	{"InvalidTopLevel3.q", &errors.InvalidTopLevel{Instruction: "+"}},
	{"MissingBlockEnd.q", errors.MissingBlockEnd},
	{"MissingBlockEnd2.q", errors.MissingBlockEnd},
	{"MissingBlockStart.q", errors.MissingBlockStart},
	{"MissingGroupEnd.q", errors.MissingGroupEnd},
	{"MissingGroupStart.q", errors.MissingGroupStart},
	{"MissingParameter.q", errors.MissingParameter},
	{"MissingParameter2.q", errors.MissingParameter},
	{"MissingParameter3.q", errors.MissingParameter},
	{"MissingType.q", errors.MissingType},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			b := build.New(filepath.Join("testdata", "errors", test.File))
			_, err := scanner.Scan(b)
			assert.NotNil(t, err)
			assert.Contains(t, err.Error(), test.ExpectedError.Error())
		})
	}
}
package scanner_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"ExpectedFunctionDefinition.q", scanner.ExpectedFunctionDefinition},
	{"ExpectedFunctionDefinition2.q", scanner.ExpectedFunctionDefinition},
	{"ExpectedPackageName.q", scanner.ExpectedPackageName},
	{"InvalidCharacter.q", &scanner.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter2.q", &scanner.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter3.q", &scanner.InvalidCharacter{Character: "@"}},
	{"InvalidCharacter4.q", &scanner.InvalidCharacter{Character: "+++"}},
	{"InvalidExpression.q", scanner.InvalidExpression},
	{"InvalidFunctionDefinition.q", scanner.InvalidFunctionDefinition},
	{"InvalidFunctionDefinition2.q", scanner.InvalidFunctionDefinition},
	{"InvalidTopLevel.q", &scanner.InvalidTopLevel{Instruction: "123"}},
	{"InvalidTopLevel2.q", &scanner.InvalidTopLevel{Instruction: "\"Hello\""}},
	{"InvalidTopLevel3.q", &scanner.InvalidTopLevel{Instruction: "+"}},
	{"MissingBlockEnd.q", scanner.MissingBlockEnd},
	{"MissingBlockEnd2.q", scanner.MissingBlockEnd},
	{"MissingBlockStart.q", scanner.MissingBlockStart},
	{"MissingExpression.q", scanner.MissingExpression},
	{"MissingGroupEnd.q", scanner.MissingGroupEnd},
	{"MissingGroupEnd2.q", scanner.MissingGroupEnd},
	{"MissingGroupStart.q", scanner.MissingGroupStart},
	{"MissingParameter.q", scanner.MissingParameter},
	{"MissingParameter2.q", scanner.MissingParameter},
	{"MissingParameter3.q", scanner.MissingParameter},
	{"MissingParameter4.q", scanner.MissingParameter},
	{"MissingParameter5.q", scanner.MissingParameter},
	{"MissingType.q", scanner.MissingType},
	{"UnknownImport.q", &scanner.UnknownImport{Package: "unknown"}},
}

func TestErrors(t *testing.T) {
	for _, test := range errs {
		name := strings.TrimSuffix(test.File, ".q")

		t.Run(name, func(t *testing.T) {
			b := config.New(filepath.Join("testdata", "errors", test.File))
			_, err := scanner.Scan(b)
			assert.NotNil(t, err)
			assert.Equal(t, err.Error(), test.ExpectedError.Error())
		})
	}
}
package compiler_test

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
	{"MissingMainFunction.q", compiler.MissingMainFunction},
	{"MultiError.q", &compiler.MultiError{Errors: []error{
		&core.UnknownIdentifier{Name: "unknown1"},
		&core.UnknownIdentifier{Name: "unknown2"},
	}}},
	{"UnusedImport.q", &compiler.UnusedImport{Package: "run"}},
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
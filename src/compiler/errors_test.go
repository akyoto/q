package compiler_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"MissingMainFunction.q", compiler.MissingMainFunction},
	{"UnusedImport.q", &compiler.UnusedImport{Package: "os"}},
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
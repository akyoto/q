package ast_test

import (
	"path/filepath"
	"strings"
	"testing"

	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

var errs = []struct {
	File          string
	ExpectedError error
}{
	{"InvalidInstruction.q", &ast.InvalidInstruction{Instruction: "()"}},
	{"InvalidInstruction2.q", &ast.InvalidInstruction{Instruction: "42"}},
	{"InvalidInstruction3.q", &ast.InvalidInstruction{Instruction: "2 + 3"}},
	{"InvalidInstruction4.q", &ast.InvalidInstruction{Instruction: "\"not used\""}},
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
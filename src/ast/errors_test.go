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
	{"EmptySwitch.q", ast.EmptySwitch},
	{"ExpectedIfBeforeElse.q", ast.ExpectedIfBeforeElse},
	{"ExpectedIfBeforeElse2.q", ast.ExpectedIfBeforeElse},
	{"InvalidInstruction.q", ast.InvalidInstruction},
	{"InvalidInstruction2.q", ast.InvalidInstruction},
	{"InvalidInstruction3.q", ast.InvalidInstruction},
	{"InvalidInstruction4.q", ast.InvalidInstruction},
	{"MissingOperand.q", ast.MissingOperand},
	{"MissingOperand2.q", ast.MissingOperand},
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
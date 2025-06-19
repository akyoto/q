package compiler_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/scanner"
	"git.urbach.dev/go/assert"
)

func TestCompile(t *testing.T) {
	b := build.New("../../examples/hello")
	_, err := compiler.Compile(scanner.Scan(b))
	assert.Nil(t, err)
}
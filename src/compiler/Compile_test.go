package compiler_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/go/assert"
)

func TestCompile(t *testing.T) {
	b := build.New("../../examples/hello")
	_, err := compiler.Compile(b)
	assert.Nil(t, err)
}
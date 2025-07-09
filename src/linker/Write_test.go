package linker_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

func TestWrite(t *testing.T) {
	b := build.New("../../examples/hello")

	b.Matrix(func(b *build.Build) {
		env, err := compiler.Compile(b)
		assert.Nil(t, err)

		writer := &exe.Discard{}
		linker.Write(writer, b, env)
	})
}
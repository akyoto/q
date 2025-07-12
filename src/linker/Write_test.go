package linker_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

func TestWrite(t *testing.T) {
	build := config.New("../../examples/hello")

	build.Matrix(func(build *config.Build) {
		env, err := compiler.Compile(build)
		assert.Nil(t, err)

		writer := &exe.Discard{}
		linker.Write(writer, env)
	})
}
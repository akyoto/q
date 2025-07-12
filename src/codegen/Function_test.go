package codegen_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/go/assert"
)

func TestHelloExample(t *testing.T) {
	build := config.New("../../examples/hello")

	build.Matrix(func(build *config.Build) {
		_, err := compiler.Compile(build)
		assert.Nil(t, err)
	})
}
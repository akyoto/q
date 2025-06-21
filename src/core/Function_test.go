package core_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	b := build.New("../../examples/hello")
	env, err := compiler.Compile(b)
	assert.Nil(t, err)

	main, exists := env.Functions["main.main"]
	assert.True(t, exists)
	assert.False(t, main.IsExtern())
	assert.Equal(t, main.UniqueName, "main.main")
	assert.Equal(t, main.String(), main.UniqueName)

	write, exists := env.Functions["io.write"]
	assert.True(t, exists)
	write.Output[0].Type()
}
package core_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/core"
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

	deps := []*core.Function{}

	main.EachDependency(map[*core.Function]bool{}, func(dep *core.Function) {
		deps = append(deps, dep)
	})

	assert.True(t, len(deps) >= 2)
	assert.Equal(t, deps[0].UniqueName, "main.main")
	assert.Equal(t, deps[1].UniqueName, "io.write")
}
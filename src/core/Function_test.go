package core_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/assert"
)

func TestFunction(t *testing.T) {
	b := config.New("../../examples/hello")
	env, err := compiler.Compile(b)
	assert.Nil(t, err)

	main := env.Function("main", "main")
	assert.NotNil(t, main)
	assert.False(t, main.IsExtern())
	assert.Equal(t, main.FullName, "main.main")
	assert.Equal(t, main.String(), main.FullName)

	deps := []*core.Function{}

	main.EachDependency(map[*core.Function]bool{}, func(dep *core.Function) {
		deps = append(deps, dep)
	})

	assert.True(t, len(deps) >= 2)
	assert.Equal(t, deps[0].FullName, "main.main")
	assert.Equal(t, deps[1].FullName, "io.write[string]")
}
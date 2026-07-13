package resolver_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/resolver"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestResolver(t *testing.T) {
	build := config.New()
	env := core.NewEnvironment(build)
	src := []byte("global { x int }")
	file := fs.NewFile("", "main", src)

	env.ReceiveGlobal(&core.Global{
		Name:   "x",
		File:   file,
		Tokens: file.Tokens[2:4],
	})

	err := resolver.ResolveTypes(env)
	assert.Nil(t, err)
	assert.Equal(t, env.Packages["main"].Globals["x"].Typ.Name(), types.Int.Name())
}
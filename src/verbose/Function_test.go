package verbose_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/verbose"
	"git.urbach.dev/go/assert"
)

func TestVerboseOutput(t *testing.T) {
	b := config.New("../../examples/hello")
	b.SetVerbose(true)
	env, err := compiler.Compile(b)
	assert.Nil(t, err)
	verbose.Show(env)
}
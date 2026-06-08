package verbose_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/verbose"
	"git.urbach.dev/go/assert"
)

func TestVerboseOutput(t *testing.T) {
	examples := "../../examples"

	fs.Walk(examples, func(dir string) {
		t.Run(dir, func(t *testing.T) {
			b := config.New(filepath.Join(examples, dir))
			env, err := compiler.Compile(b)
			assert.Nil(t, err)
			verbose.ASM(env.Init)
			verbose.SSA(env.Init)
		})
	})
}
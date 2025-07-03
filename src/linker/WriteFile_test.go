package linker_test

import (
	"os"
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/compiler"
	"git.urbach.dev/cli/q/src/linker"
	"git.urbach.dev/go/assert"
)

func TestWriteFile(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "q", "tests")
	err := os.MkdirAll(tmpDir, 0o755)
	assert.Nil(t, err)

	fromPath := "../../examples/hello/hello.q"
	contents, err := os.ReadFile(fromPath)
	assert.Nil(t, err)

	toPath := filepath.Join(tmpDir, "hello.q")
	err = os.WriteFile(toPath, contents, 0o755)
	assert.Nil(t, err)

	b := build.New(toPath)
	env, err := compiler.Compile(b)
	assert.Nil(t, err)

	b.Arch = build.ARM
	err = linker.WriteFile(b.Executable(), b, env)
	assert.Nil(t, err)

	b.Arch = build.X86
	err = linker.WriteFile(b.Executable(), b, env)
	assert.Nil(t, err)
}
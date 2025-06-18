package build_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/go/assert"
)

func TestExecutableNameFromDirectory(t *testing.T) {
	b := build.New("../../examples/hello")
	exe := filepath.Base(b.Executable())

	if b.OS == build.Windows {
		assert.Equal(t, exe, "hello.exe")
	} else {
		assert.Equal(t, exe, "hello")
	}
}

func TestExecutableNameFromFile(t *testing.T) {
	b := build.New("../../examples/hello/hello.q")
	exe := filepath.Base(b.Executable())

	if b.OS == build.Windows {
		assert.Equal(t, exe, "hello.exe")
	} else {
		assert.Equal(t, exe, "hello")
	}
}
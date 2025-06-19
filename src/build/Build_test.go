package build_test

import (
	"path/filepath"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/go/assert"
)

func TestExecutableFromDirectoryLinux(t *testing.T) {
	b := build.New("../../examples/hello")
	b.OS = build.Linux
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello")
}

func TestExecutableFromFileLinux(t *testing.T) {
	b := build.New("../../examples/hello/hello.q")
	b.OS = build.Linux
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello")
}

func TestExecutableFromDirectoryMac(t *testing.T) {
	b := build.New("../../examples/hello")
	b.OS = build.Mac
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello")
}

func TestExecutableFromFileMac(t *testing.T) {
	b := build.New("../../examples/hello/hello.q")
	b.OS = build.Mac
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello")
}

func TestExecutableFromDirectoryWindows(t *testing.T) {
	b := build.New("../../examples/hello")
	b.OS = build.Windows
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello.exe")
}

func TestExecutableFromFileWindows(t *testing.T) {
	b := build.New("../../examples/hello/hello.q")
	b.OS = build.Windows
	exe := filepath.Base(b.Executable())
	assert.Equal(t, exe, "hello.exe")
}
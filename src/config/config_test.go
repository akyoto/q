package config_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/go/assert"
)

func TestNew(t *testing.T) {
	osList := []string{"linux", "darwin", "windows"}
	archList := []string{"amd64", "arm64"}
	fileList := []string{"../../examples/hello", "../../examples/hello/hello.q"}

	for _, os := range osList {
		for _, arch := range archList {
			global.OS = os
			global.Arch = arch

			t.Run(fmt.Sprintf("%s-%s", os, arch), func(t *testing.T) {
				for _, file := range fileList {
					b := config.New(file)
					assert.NotEqual(t, b.Arch, config.UnknownArch)
					assert.NotEqual(t, b.OS, config.UnknownOS)
				}
			})
		}
	}

	global.OS = runtime.GOOS
	global.Arch = runtime.GOARCH
}

func TestFields(t *testing.T) {
	b := config.New("../../examples/hello")

	b.Matrix(func(b *config.Build) {
		assert.NotEqual(t, b.Arch, config.UnknownArch)
		assert.NotEqual(t, b.OS, config.UnknownOS)

		exe := filepath.Base(b.Executable())

		if b.OS == config.Windows {
			assert.Equal(t, exe, "hello.exe")
		} else {
			assert.Equal(t, exe, "hello")
		}

		assert.True(t, b.FileAlign() > 0)
		assert.True(t, b.FileAlign() <= b.MemoryAlign())

		if b.OS == config.Linux {
			assert.True(t, b.Congruent())
		}

		assert.NotNil(t, b.CPU())
	})
}
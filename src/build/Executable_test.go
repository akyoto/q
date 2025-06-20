package build_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"git.urbach.dev/cli/q/src/build"
	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/go/assert"
)

func TestExecutable(t *testing.T) {
	osList := []string{"linux", "darwin", "windows"}
	archList := []string{"amd64", "arm64"}
	fileList := []string{"../../examples/hello", "../../examples/hello/hello.q"}

	for _, os := range osList {
		for _, arch := range archList {
			t.Run(fmt.Sprintf("%s-%s", os, arch), func(t *testing.T) {
				for _, file := range fileList {
					global.OS = os
					global.Arch = arch
					b := build.New(file)
					exe := filepath.Base(b.Executable())

					if os == "windows" {
						assert.Equal(t, exe, "hello.exe")
					} else {
						assert.Equal(t, exe, "hello")
					}
				}
			})
		}
	}

	global.OS = runtime.GOOS
	global.Arch = runtime.GOARCH
}
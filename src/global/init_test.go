package global_test

import (
	"runtime"
	"testing"

	"git.urbach.dev/cli/q/src/global"
	"git.urbach.dev/go/assert"
)

func TestInit(t *testing.T) {
	assert.Equal(t, global.Arch, runtime.GOARCH)
	assert.True(t, len(global.Executable) > 0)
	assert.True(t, len(global.Library) > 0)
	assert.Equal(t, global.OS, runtime.GOOS)
	assert.True(t, len(global.Root) > 0)
}
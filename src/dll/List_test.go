package dll_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/go/assert"
)

func TestList(t *testing.T) {
	libs := dll.List{}
	assert.False(t, libs.Contains("kernel32"))
	assert.Equal(t, -1, libs.Index("kernel32", "ExitProcess"))

	libs.Append("kernel32", "ExitProcess")
	assert.True(t, libs.Contains("kernel32"))
	assert.Equal(t, 0, libs.Index("kernel32", "ExitProcess"))

	libs.Append("user32", "MessageBox")
	assert.True(t, libs.Contains("kernel32"))
	assert.True(t, libs.Contains("user32"))
	assert.Equal(t, 0, libs.Index("kernel32", "ExitProcess"))
	assert.Equal(t, 2, libs.Index("user32", "MessageBox"))

	libs.Append("kernel32", "SetConsoleCP")
	assert.Equal(t, 0, libs.Index("kernel32", "ExitProcess"))
	assert.Equal(t, 1, libs.Index("kernel32", "SetConsoleCP"))
	assert.Equal(t, 3, libs.Index("user32", "MessageBox"))

	libs.Append("kernel32", "ExitProcess")
	assert.Equal(t, 0, libs.Index("kernel32", "ExitProcess"))
	assert.Equal(t, 1, libs.Index("kernel32", "SetConsoleCP"))
	assert.Equal(t, 3, libs.Index("user32", "MessageBox"))
}
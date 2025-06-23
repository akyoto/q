package exe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/go/assert"
)

func TestExecutable(t *testing.T) {
	align := 0x1000
	x := exe.New(1, align, align)
	x.InitSections([]byte{1}, []byte{1})
	assert.Equal(t, len(x.Sections), 2)
	assert.Equal(t, x.Sections[0].Padding, align-1)
	assert.Equal(t, x.Sections[0].FileOffset, align)
	assert.Equal(t, x.Sections[1].Padding, align-1)
	assert.Equal(t, x.Sections[1].FileOffset, align*2)
}
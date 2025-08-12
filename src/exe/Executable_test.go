package exe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/go/assert"
)

func TestSimple(t *testing.T) {
	align := 32
	x := exe.New(1, align, align, false, false, []byte{1}, []byte{1})
	assert.Equal(t, len(x.Sections), 2)
	assert.Equal(t, x.Sections[0].Padding, align-1)
	assert.Equal(t, x.Sections[0].FileOffset, align)
	assert.Equal(t, x.Sections[1].Padding, align-1)
	assert.Equal(t, x.Sections[1].FileOffset, align*2)
}

func TestCongruent(t *testing.T) {
	fileAlign := 16
	memoryAlign := 32
	x := exe.New(1, fileAlign, memoryAlign, true, false, []byte{1}, []byte{1}, []byte{1})
	assert.Equal(t, len(x.Sections), 3)
	assert.Equal(t, x.Sections[0].FileOffset, fileAlign)
	assert.Equal(t, x.Sections[1].FileOffset, fileAlign*2)
	assert.Equal(t, x.Sections[2].FileOffset, fileAlign*3)
	assert.Equal(t, x.Sections[0].MemoryOffset, memoryAlign+(fileAlign%memoryAlign))
	assert.Equal(t, x.Sections[1].MemoryOffset, memoryAlign*2+(fileAlign*2%memoryAlign))
	assert.Equal(t, x.Sections[2].MemoryOffset, memoryAlign*3+(fileAlign*3%memoryAlign))
}
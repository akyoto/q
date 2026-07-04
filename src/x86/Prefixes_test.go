package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestPrefix(t *testing.T) {
	assert.DeepEqual(t, x86.Lock(nil), []byte{0xF0})
	assert.DeepEqual(t, x86.SegmentBaseFS(nil), []byte{0x64})
	assert.DeepEqual(t, x86.SegmentBaseGS(nil), []byte{0x65})
}
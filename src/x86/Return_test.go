package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestReturn(t *testing.T) {
	assert.DeepEqual(t, x86.Return(nil), []byte{0xC3})
}
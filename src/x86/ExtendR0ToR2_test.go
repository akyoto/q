package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestExtendR0ToR2(t *testing.T) {
	assert.DeepEqual(t, x86.ExtendR0ToR2(nil), []byte{0x48, 0x99})
}
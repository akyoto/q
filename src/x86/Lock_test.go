package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestLock(t *testing.T) {
	assert.DeepEqual(t, x86.Lock(nil), []byte{0xF0})
}
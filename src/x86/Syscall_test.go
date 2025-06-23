package x86_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/x86"
	"git.urbach.dev/go/assert"
)

func TestSyscall(t *testing.T) {
	assert.DeepEqual(t, x86.Syscall(nil), []byte{0x0F, 0x05})
}
package cpu_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestRegisterString(t *testing.T) {
	r1 := cpu.Register(1)
	assert.Equal(t, r1.String(), "r1")

	undefined := cpu.Register(-1)
	assert.Equal(t, undefined.String(), "r?")
}
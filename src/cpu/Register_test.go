package cpu_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestRegisterString(t *testing.T) {
	register := cpu.Register(1)
	assert.Equal(t, "r1", register.String())
}
package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/go/assert"
)

func TestConstants(t *testing.T) {
	assert.DeepEqual(t, arm.Nop(), 0xD503201F)
	assert.DeepEqual(t, arm.Return(), 0xD65F03C0)
	assert.DeepEqual(t, arm.Syscall(), 0xD4000001)
}

func TestNotEncodable(t *testing.T) {
	_, encodable := arm.AndRegisterNumber(arm.X0, arm.X0, 0)
	assert.False(t, encodable)
	_, encodable = arm.OrRegisterNumber(arm.X0, arm.X0, 0)
	assert.False(t, encodable)
	_, encodable = arm.XorRegisterNumber(arm.X0, arm.X0, 0)
	assert.False(t, encodable)
	_, encodable = arm.AndRegisterNumber(arm.X0, arm.X0, -1)
	assert.False(t, encodable)
	_, encodable = arm.OrRegisterNumber(arm.X0, arm.X0, -1)
	assert.False(t, encodable)
	_, encodable = arm.XorRegisterNumber(arm.X0, arm.X0, -1)
	assert.False(t, encodable)
	_, encodable = arm.AddRegisterNumber(arm.X0, arm.X0, 0xFFFF)
	assert.False(t, encodable)
	_, encodable = arm.AddRegisterNumber(arm.X0, arm.X0, 0xF0000000)
	assert.False(t, encodable)
	_, encodable = arm.SubRegisterNumber(arm.X0, arm.X0, 0xFFFF)
	assert.False(t, encodable)
	_, encodable = arm.SubRegisterNumber(arm.X0, arm.X0, 0xF0000000)
	assert.False(t, encodable)
}
package arm_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/arm"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestMoveRegisterRegister(t *testing.T) {
	usagePatterns := []struct {
		Destination cpu.Register
		Source      cpu.Register
		Code        uint32
	}{
		{arm.X0, arm.X1, 0xAA0103E0},
		{arm.X1, arm.X0, 0xAA0003E1},
		{arm.FP, arm.SP, 0x910003FD},
		{arm.SP, arm.FP, 0x910003BF},
	}

	for _, pattern := range usagePatterns {
		t.Logf("mov %s, %s", pattern.Destination, pattern.Source)
		code := arm.MoveRegisterRegister(pattern.Destination, pattern.Source)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestMoveRegisterNumber(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Number   uint64
		Code     []byte
	}{
		{arm.X0, 0, []byte{0x00, 0x00, 0x80, 0xD2}},
		{arm.X0, 0xCAFEBABE, []byte{0xC0, 0x57, 0x97, 0xD2, 0xC0, 0x5F, 0xB9, 0xF2}},
		{arm.X0, 0xDEADC0DE, []byte{0xC0, 0x1B, 0x98, 0xD2, 0xA0, 0xD5, 0xBB, 0xF2}},
	}

	for _, pattern := range usagePatterns {
		t.Logf("mov %s, 0x%X", pattern.Register, pattern.Number)
		code := arm.MoveRegisterNumber(nil, pattern.Register, int(pattern.Number))
		assert.DeepEqual(t, code, pattern.Code)
	}
}

func TestMoveRegisterNumberSI(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Number   uint64
		Code     uint32
	}{
		// MOVZ
		{arm.X0, 0x0, 0xD2800000},
		{arm.X0, 0x1, 0xD2800020},
		{arm.X0, 0x1000, 0xD2820000},

		// MOV (bitmask immediate)
		{arm.X0, 0x1FFFF, 0xB24043E0},
		{arm.X0, 0x7FFFFFFF, 0xB2407BE0},
		{arm.X0, 0xFFFFFFFF, 0xB2407FE0},
		{arm.X0, 0xC3FFFFFFC3FFFFFF, 0xB2026FE0},

		// MOV (inverted wide immediate)
		{arm.X0, 0xFFFFFFFFFFFFFFFF, 0x92800000},
		{arm.X0, 0x7FFFFFFFFFFFFFFF, 0x92F00000},
		{arm.X0, 0x2FFFFFFFF, 0x92DFFFA0}, // not encodable in the GNU assembler
		{arm.X0, 0x2FFFF, 0x92BFFFA0},     // not encodable in the GNU assembler

		// Not encodable
		{arm.X0, 0xCAFEBABE, 0},
		{arm.X0, 0xDEADC0DE, 0},
	}

	for _, pattern := range usagePatterns {
		t.Logf("mov %s, %d", pattern.Register, pattern.Number)
		code, encodable := arm.MoveRegisterNumberSI(pattern.Register, int(pattern.Number))

		if pattern.Code != 0 {
			assert.True(t, encodable)
			assert.Equal(t, code, pattern.Code)
		} else {
			assert.False(t, encodable)
		}
	}
}

func TestMoveKeep(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Number   uint16
		Code     uint32
	}{
		{arm.X0, 0, 0xF2800000},
		{arm.X0, 1, 0xF2800020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("movk %s, %d", pattern.Register, pattern.Number)
		code := arm.MoveKeep(pattern.Register, 0, pattern.Number)
		assert.Equal(t, code, pattern.Code)
	}
}

func TestMoveZero(t *testing.T) {
	usagePatterns := []struct {
		Register cpu.Register
		Number   uint16
		Code     uint32
	}{
		{arm.X0, 0, 0xD2800000},
		{arm.X0, 1, 0xD2800020},
	}

	for _, pattern := range usagePatterns {
		t.Logf("movz %s, %d", pattern.Register, pattern.Number)
		code := arm.MoveZero(pattern.Register, 0, pattern.Number)
		assert.Equal(t, code, pattern.Code)
	}
}
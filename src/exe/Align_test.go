package exe_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/go/assert"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		Address   uint
		Alignment uint
		Expected  uint
	}{
		{0, 16, 0},
		{1, 16, 16},
		{16, 16, 16},
		{17, 16, 32},
		{100, 64, 128},
	}

	for _, test := range tests {
		aligned := exe.Align(test.Address, test.Alignment)
		assert.Equal(t, aligned, test.Expected)
	}
}

func TestAlignPad(t *testing.T) {
	tests := []struct {
		Address         uint
		Alignment       uint
		ExpectedAddress uint
		ExpectedPadding uint
	}{
		{0, 16, 0, 0},
		{1, 16, 16, 15},
		{15, 16, 16, 1},
		{16, 16, 16, 0},
		{17, 16, 32, 15},
		{63, 32, 64, 1},
	}

	for _, test := range tests {
		aligned, padding := exe.AlignPad(test.Address, test.Alignment)
		assert.Equal(t, aligned, test.ExpectedAddress)
		assert.Equal(t, padding, test.ExpectedPadding)
	}
}

func TestPad(t *testing.T) {
	tests := []struct {
		Address   uint
		Alignment uint
		Expected  uint
	}{
		{0, 16, 0},
		{1, 16, 15},
		{15, 16, 1},
		{16, 16, 0},
		{17, 16, 15},
		{100, 64, 28},
	}

	for _, test := range tests {
		padding := exe.Pad(test.Address, test.Alignment)
		assert.Equal(t, padding, test.Expected)
	}
}

func TestPadSlice(t *testing.T) {
	tests := []struct {
		Input          []byte
		Alignment      uint
		ExpectedOutput []byte
	}{
		{nil, 4, nil},
		{[]byte{1, 2, 3, 4}, 4, []byte{1, 2, 3, 4}},
		{[]byte{1, 2, 3}, 4, []byte{1, 2, 3, 0}},
		{[]byte{1}, 8, []byte{1, 0, 0, 0, 0, 0, 0, 0}},
		{make([]byte, 3, 8), 8, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, test := range tests {
		padded := exe.PadSlice(test.Input, test.Alignment)
		assert.DeepEqual(t, padded, test.ExpectedOutput)
	}
}
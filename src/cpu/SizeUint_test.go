package cpu_test

import (
	"math"
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestSizeUint(t *testing.T) {
	assert.Equal(t, cpu.SizeUint(0), 1)
	assert.Equal(t, cpu.SizeUint(math.MaxUint8), 1)
	assert.Equal(t, cpu.SizeUint(math.MaxUint8+1), 2)
	assert.Equal(t, cpu.SizeUint(math.MaxUint16), 2)
	assert.Equal(t, cpu.SizeUint(math.MaxUint16+1), 4)
	assert.Equal(t, cpu.SizeUint(math.MaxUint32), 4)
	assert.Equal(t, cpu.SizeUint(math.MaxUint32+1), 8)
}
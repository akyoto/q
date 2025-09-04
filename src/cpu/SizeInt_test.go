package cpu_test

import (
	"math"
	"testing"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

func TestSizeInt(t *testing.T) {
	assert.Equal(t, cpu.SizeInt(0), 1)
	assert.Equal(t, cpu.SizeInt(math.MinInt8), 1)
	assert.Equal(t, cpu.SizeInt(math.MaxInt8), 1)
	assert.Equal(t, cpu.SizeInt(math.MinInt16), 2)
	assert.Equal(t, cpu.SizeInt(math.MaxInt16), 2)
	assert.Equal(t, cpu.SizeInt(math.MinInt32), 4)
	assert.Equal(t, cpu.SizeInt(math.MaxInt32), 4)
	assert.Equal(t, cpu.SizeInt(math.MinInt64), 8)
	assert.Equal(t, cpu.SizeInt(math.MaxInt64), 8)
}
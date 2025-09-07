package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestName(t *testing.T) {
	assert.Equal(t, types.Int.Name(), "int64")
	assert.Equal(t, types.AnyPointer.Name(), "*any")
	assert.Equal(t, (&types.Pointer{To: types.Int}).Name(), "*int64")
	assert.Equal(t, types.String.Name(), "string")
}

func TestSize(t *testing.T) {
	assert.Equal(t, types.Int.Size(), 8)
	assert.Equal(t, types.Int8.Size(), 1)
	assert.Equal(t, types.Int16.Size(), 2)
	assert.Equal(t, types.Int32.Size(), 4)
	assert.Equal(t, types.Int64.Size(), 8)
	assert.Equal(t, types.AnyPointer.Size(), 8)
	assert.Equal(t, types.String.Size(), 16)
	assert.Equal(t, (&types.Pointer{To: types.Int}).Size(), 8)
}

func TestBasics(t *testing.T) {
	assert.True(t, types.Is(types.Int, types.Int))
	assert.True(t, types.Is(types.Int, types.Any))
	assert.True(t, types.Is(types.Int, types.AnyInt))
	assert.True(t, types.Is(types.AnyInt, types.AnyInt))
	assert.True(t, types.Is(types.AnyInt, types.Int))
	assert.True(t, types.Is(types.AnyPointer, types.AnyPointer))
	assert.False(t, types.Is(types.Int, types.Float))
	assert.False(t, types.Is(types.AnyPointer, types.AnyInt))
	assert.False(t, types.Is(&types.Pointer{To: types.Int}, &types.Pointer{To: types.Float}))
	assert.False(t, types.Is(types.Any, types.Int))
	assert.False(t, types.Is(types.Any, types.Float))

	// TODO: This check is currently disabled due to some temporary hacks, add it back later.
	//assert.False(t, types.Is(types.AnyInt, types.AnyPointer))
}

func TestUnions(t *testing.T) {
	intOrFloat := &types.Union{Types: []types.Type{types.Int, types.Float}}
	assert.True(t, types.Is(types.Int, intOrFloat))
	assert.True(t, types.Is(types.Float, intOrFloat))
	assert.False(t, types.Is(types.String, intOrFloat))
	assert.Equal(t, intOrFloat.Index(types.Int), 0)
	assert.Equal(t, intOrFloat.Index(types.Float), 1)
}

func TestSpecialCases(t *testing.T) {
	assert.True(t, types.Is(&types.Pointer{To: types.Int}, types.AnyPointer))
	assert.True(t, types.Is(&types.Pointer{To: types.Float}, types.AnyPointer))
}
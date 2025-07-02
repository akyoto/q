package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestName(t *testing.T) {
	assert.Equal(t, types.Int.Name(), "int64")
	assert.Equal(t, types.AnyArray.Name(), "[]any")
	assert.Equal(t, types.AnyPointer.Name(), "*any")
	assert.Equal(t, (&types.Pointer{To: types.Int}).Name(), "*int64")
	assert.Equal(t, (&types.Array{Of: types.Int}).Name(), "[]int64")
	assert.Equal(t, types.String.Name(), "string")
}

func TestSize(t *testing.T) {
	assert.Equal(t, types.Int.Size(), 8)
	assert.Equal(t, types.Int8.Size(), 1)
	assert.Equal(t, types.Int16.Size(), 2)
	assert.Equal(t, types.Int32.Size(), 4)
	assert.Equal(t, types.Int64.Size(), 8)
	assert.Equal(t, types.AnyArray.Size(), 8)
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
	assert.True(t, types.Is(&types.Array{Of: types.Int}, types.AnyArray))
	assert.False(t, types.Is(types.Int, types.Float))
	assert.False(t, types.Is(types.Any, types.Int))
	assert.False(t, types.Is(types.AnyPointer, types.AnyInt))
	assert.False(t, types.Is(types.AnyInt, types.AnyPointer))
	assert.False(t, types.Is(&types.Pointer{To: types.Int}, &types.Pointer{To: types.Float}))
}

func TestSpecialCases(t *testing.T) {
	// Case #1:
	// For syscalls whose return type is `nil` we currently allow casting them to anything.
	assert.True(t, types.Is(nil, types.Int))
	assert.True(t, types.Is(nil, types.Float))

	// Case #2:
	// A pointer pointing to a known type fulfills the requirement of a pointer to anything.
	assert.True(t, types.Is(&types.Pointer{To: types.Int}, types.AnyPointer))
	assert.True(t, types.Is(&types.Pointer{To: types.Float}, types.AnyPointer))

	// Case #3:
	// Arrays are also just pointers.
	assert.True(t, types.Is(&types.Array{Of: types.Int}, types.AnyPointer))
	assert.True(t, types.Is(&types.Array{Of: types.Float}, types.AnyPointer))
}
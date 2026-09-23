package types_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestStruct(t *testing.T) {
	point := types.NewStruct(nil, "math", "Point")

	x := &types.Field{
		Name:   "x",
		Type:   types.Int,
		Index:  0,
		Offset: 0,
	}

	y := &types.Field{
		Name:   "y",
		Type:   types.Int,
		Index:  1,
		Offset: 8,
	}

	point.AddField(x)
	point.AddField(y)

	assert.Equal(t, x.String(), "x")
	assert.Equal(t, y.String(), "y")
	assert.Equal(t, point.FieldByName("x"), x)
	assert.Equal(t, point.FieldByName("y"), y)
	assert.Nil(t, point.FieldByName("invalid"))
}

func TestArray(t *testing.T) {
	array := types.Array(types.UInt8, 20)
	assert.Equal(t, array.Name(), "[20]uint8")
	assert.Equal(t, array.Size(), 20)
	assert.Equal(t, len(array.Fields), 20)
	assert.Equal(t, array.Fields[0].Offset, uint64(0))
	assert.Equal(t, array.Fields[19].Offset, uint64(19))

	array = types.Array(types.Int64, 4)
	assert.Equal(t, array.Name(), "[4]int64")
	assert.Equal(t, array.Size(), 32)
	assert.Equal(t, len(array.Fields), 4)
	assert.Equal(t, array.Fields[0].Offset, uint64(0))
	assert.Equal(t, array.Fields[3].Offset, uint64(24))
}
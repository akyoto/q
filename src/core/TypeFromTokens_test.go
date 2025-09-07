package core_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
	"git.urbach.dev/go/assert"
)

func TestTypeFromTokens(t *testing.T) {
	tests := []struct {
		Source       string
		ExpectedType types.Type
	}{
		{"int", types.Int},
		{"int64", types.Int64},
		{"int32", types.Int32},
		{"int16", types.Int16},
		{"int8", types.Int8},
		{"uint", types.UInt},
		{"uint64", types.UInt64},
		{"uint32", types.UInt32},
		{"uint16", types.UInt16},
		{"uint8", types.UInt8},
		{"byte", types.Byte},
		{"bool", types.Bool},
		{"float", types.Float},
		{"float64", types.Float64},
		{"float32", types.Float32},
		{"string", types.String},
		{"any", types.Any},
		{"*any", types.AnyPointer},
		{"*byte", &types.Pointer{To: types.Byte}},
		{"123", nil},
		{"*", nil},
		{"[]", nil},
		{"[", nil},
		{"_", nil},
	}

	for _, test := range tests {
		t.Run(test.Source, func(t *testing.T) {
			src := []byte(test.Source)
			tokens := token.Tokenize(src)
			file := &fs.File{Tokens: tokens, Bytes: src}
			typ, _ := core.TypeFromTokens(tokens, file, nil)
			assert.True(t, types.Is(typ, test.ExpectedType))
		})
	}
}
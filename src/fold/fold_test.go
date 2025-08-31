package fold_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/fold"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/go/assert"
)

func TestFold(t *testing.T) {
	operations := []token.Kind{
		token.Add,
		token.And,
		token.Div,
		token.Mul,
		token.Mod,
		token.Or,
		token.Shl,
		token.Shr,
		token.Sub,
		token.Xor,
	}

	for _, op := range operations {
		ir := ssa.IR{}
		block := ssa.NewBlock("test")
		ir.AddBlock(block)
		one := ir.Append(&ssa.Int{Int: 1})
		two := ir.Append(&ssa.Int{Int: 2})
		binOp := ir.Append(&ssa.BinaryOp{Op: op, Left: one, Right: two})

		folded := fold.Constants(ir)
		assert.NotNil(t, folded)
		assert.Equal(t, block.Index(binOp), -1)

		_, oneFolded := folded[one]
		assert.True(t, oneFolded)

		_, twoFolded := folded[two]
		assert.True(t, twoFolded)
	}
}
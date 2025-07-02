package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

// IR is a list of basic blocks.
type IR struct {
	Blocks []*Block
	nextId int
}

// AddBlock adds a new block to the function.
func (f *IR) AddBlock() *Block {
	block := &Block{
		Instructions: make([]Value, 0, 8),
	}

	f.Blocks = append(f.Blocks, block)
	return block
}

// Append adds a new value to the last block.
func (f *IR) Append(instr Value) Value {
	if len(f.Blocks) == 0 {
		f.AddBlock()
	}

	if instr.IsConst() {
		for existing := range f.Values {
			if existing.IsConst() && instr.Equals(existing) {
				return existing
			}
		}
	}

	instr.SetID(f.nextId)
	f.nextId++
	return f.Blocks[len(f.Blocks)-1].Append(instr)
}

// AppendInt adds a new integer value to the last block.
func (f *IR) AppendInt(x int) Value {
	return f.Append(&Int{Int: x})
}

// AppendFunction adds a new function value to the last block.
func (f *IR) AppendFunction(name string, typ *types.Function, extern bool) Value {
	return f.Append(&Function{UniqueName: name, Typ: typ, IsExtern: extern})
}

// AppendBytes adds a new byte slice value to the last block.
func (f *IR) AppendBytes(s []byte) Value {
	return f.Append(&Bytes{Bytes: s})
}

// AppendString adds a new string value to the last block.
func (f *IR) AppendString(s string) Value {
	return f.Append(&Bytes{Bytes: []byte(s)})
}

// Values yields on each value.
func (f *IR) Values(yield func(Value) bool) {
	for _, block := range f.Blocks {
		for _, instr := range block.Instructions {
			if !yield(instr) {
				return
			}
		}
	}
}
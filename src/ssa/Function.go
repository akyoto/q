package ssa

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// Function is a list of basic blocks.
type Function struct {
	Blocks []*Block
}

// AddBlock adds a new block to the function.
func (f *Function) AddBlock() *Block {
	block := &Block{}
	f.Blocks = append(f.Blocks, block)
	return block
}

// Append adds a new value to the last block.
func (f *Function) Append(instr Value) *Value {
	if len(f.Blocks) == 0 {
		f.Blocks = append(f.Blocks, &Block{})
	}

	if instr.IsConst() {
		for _, b := range f.Blocks {
			for _, existing := range b.Instructions {
				if instr.Equals(existing) {
					return &existing
				}
			}
		}
	}

	return f.Blocks[len(f.Blocks)-1].Append(instr)
}

// AppendInt adds a new integer value to the last block.
func (f *Function) AppendInt(x int) *Value {
	return f.Append(Value{Type: Int, Int: x})
}

// AppendRegister adds a new register value to the last block.
func (f *Function) AppendRegister(reg cpu.Register) *Value {
	return f.Append(Value{Type: Register, Register: reg})
}

// AppendFunction adds a new function value to the last block.
func (f *Function) AppendFunction(name string) *Value {
	return f.Append(Value{Type: Func, Text: name})
}

// AppendBytes adds a new byte slice value to the last block.
func (f *Function) AppendBytes(s []byte) *Value {
	return f.Append(Value{Type: String, Text: string(s)})
}

// AppendString adds a new string value to the last block.
func (f *Function) AppendString(s string) *Value {
	return f.Append(Value{Type: String, Text: s})
}
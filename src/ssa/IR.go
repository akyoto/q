package ssa

// IR is a list of basic blocks.
type IR struct {
	Blocks []*Block
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

	return f.Blocks[len(f.Blocks)-1].Append(instr)
}

// AppendInt adds a new integer value to the last block.
func (f *IR) AppendInt(x int) *Int {
	v := &Int{Int: x}
	f.Append(v)
	return v
}

// AppendRegister adds a new register value to the last block.
func (f *IR) AppendRegister(index int) *Parameter {
	v := &Parameter{Index: uint8(index)}
	f.Append(v)
	return v
}

// AppendFunction adds a new function value to the last block.
func (f *IR) AppendFunction(name string) *Function {
	v := &Function{UniqueName: name}
	f.Append(v)
	return v
}

// AppendBytes adds a new byte slice value to the last block.
func (f *IR) AppendBytes(s []byte) *Bytes {
	v := &Bytes{Bytes: s}
	f.Append(v)
	return v
}

// AppendString adds a new string value to the last block.
func (f *IR) AppendString(s string) *Bytes {
	v := &Bytes{Bytes: []byte(s)}
	f.Append(v)
	return v
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
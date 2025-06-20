package ssa

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
package ssa

// IR is a list of basic blocks.
type IR struct {
	Blocks []*Block
}

// AddBlock adds a new block to the function.
func (f *IR) AddBlock(block *Block) {
	f.Blocks = append(f.Blocks, block)
}

// Append adds a new value to the last block.
func (f *IR) Append(instr Value) Value {
	existing := f.Block().FindExisting(instr)

	if existing != nil {
		return existing
	}

	f.Block().Append(instr)
	return instr
}

// Block returns the last block.
func (f *IR) Block() *Block {
	return f.Blocks[len(f.Blocks)-1]
}

// CountValues returns the total number of values.
func (f *IR) CountValues() int {
	count := 0

	for _, block := range f.Blocks {
		count += len(block.Instructions)
	}

	return count
}

// Finalize creates the list of users for each value.
func (f *IR) Finalize() {
	for _, value := range f.Values {
		for _, input := range value.Inputs() {
			input.addUser(value)
		}
	}
}

// Values yields on each value.
func (f *IR) Values(yield func(int, Value) bool) {
	index := 0

	for _, block := range f.Blocks {
		for _, instr := range block.Instructions {
			if !yield(index, instr) {
				return
			}

			index++
		}
	}
}
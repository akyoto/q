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
	existing := f.FindExisting(instr)

	if existing != nil {
		return existing
	}

	return f.Blocks[len(f.Blocks)-1].Append(instr)
}

// CountValues returns the total number of values.
func (f *IR) CountValues() int {
	count := 0

	for _, block := range f.Blocks {
		count += len(block.Instructions)
	}

	return count
}

// FindExisting returns an equal instruction that's already appended or `nil` if none could be found.
func (f *IR) FindExisting(instr Value) Value {
	if !instr.IsConst() {
		return nil
	}

	for _, existing := range f.Values {
		if existing.IsConst() && instr.Equals(existing) {
			return existing
		}
	}

	return nil
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
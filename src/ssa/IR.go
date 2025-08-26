package ssa

// IR is a list of basic blocks.
type IR struct {
	Blocks []*Block
}

// AddBlock adds a new block to the function.
func (ir *IR) AddBlock(block *Block) {
	ir.Blocks = append(ir.Blocks, block)
}

// Append adds a new value to the last block.
func (ir *IR) Append(instr Value) Value {
	existing := ir.Block().FindExisting(instr)

	if existing != nil {
		return existing
	}

	ir.Block().Append(instr)
	return instr
}

// Block returns the last block.
func (ir *IR) Block() *Block {
	return ir.Blocks[len(ir.Blocks)-1]
}

// CountValues returns the total number of values.
func (ir *IR) CountValues() int {
	count := 0

	for _, block := range ir.Blocks {
		count += len(block.Instructions)
	}

	return count
}

// ExitBlocks is an iterator for all exit blocks.
func (ir *IR) ExitBlocks(yield func(*Block) bool) {
	for i, block := range ir.Blocks {
		if i < len(ir.Blocks)-1 {
			last := block.Last()
			_, isReturn := last.(*Return)

			if !isReturn {
				continue
			}
		}

		if !yield(block) {
			return
		}
	}
}

// Finalize creates the list of users for each value.
func (ir *IR) Finalize() {
	for _, block := range ir.Blocks {
		for _, value := range block.Instructions {
			for _, input := range value.Inputs() {
				input.AddUser(value)
			}
		}
	}
}

// IsIdentified returns true if the value can be obtained from one of the identifiers.
func (ir *IR) IsIdentified(value Value) bool {
	for _, block := range ir.Blocks {
		if block.IsIdentified(value) {
			return true
		}
	}

	return false
}
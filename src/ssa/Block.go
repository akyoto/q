package ssa

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Instructions []Value
	Label        string
}

// Append adds a new instruction to the block.
func (block *Block) Append(instr Value) Value {
	for _, dep := range instr.Inputs() {
		dep.(HasLiveness).AddUser(instr)
	}

	block.Instructions = append(block.Instructions, instr)
	return instr
}
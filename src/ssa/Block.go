package ssa

type Block struct {
	Entry        []*Block
	Instructions []Instruction
	Exit         []*Block
}

func (b *Block) Append(instr Instruction) {
	b.Instructions = append(b.Instructions, instr)
}
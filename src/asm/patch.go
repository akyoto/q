package asm

// patch is a modification of a machine code instruction.
type patch struct {
	start int
	end   int
	apply func([]byte) []byte
}
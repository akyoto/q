package asm

// patch is a modification of a machine code instruction.
type patch struct {
	apply func([]byte) []byte
	start int
	end   int
}
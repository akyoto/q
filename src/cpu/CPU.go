package cpu

// CPU represents the processor.
type CPU struct {
	General           []Register
	Call              ABI
	ExternCall        ABI
	Syscall           ABI
	CasClobbered      []Register
	DivisionClobbered []Register
	DivisorRestricted []Register
	ShiftClobbered    []Register
	ShiftRestricted   []Register
	FramePointer      Register
	StackPointer      Register
	MaxRegisters      Register
}
package cpu

// CPU represents the processor.
type CPU struct {
	Call       []Register
	Syscall    []Register
	ExternCall []Register
	Return     []Register
}
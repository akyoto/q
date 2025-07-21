package cpu

// CPU represents the processor.
type CPU struct {
	General    []Register
	Division   []Register
	Shift      []Register
	Call       ABI
	ExternCall ABI
	Syscall    ABI
}
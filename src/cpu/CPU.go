package cpu

// CPU represents the processor.
type CPU struct {
	General    []Register
	Call       ABI
	ExternCall ABI
	Syscall    ABI
}
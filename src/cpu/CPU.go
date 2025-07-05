package cpu

// CPU represents the processor.
type CPU struct {
	Call       ABI
	ExternCall ABI
	Syscall    ABI
}
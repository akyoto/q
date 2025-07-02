package cpu

// CPU represents the processor.
type CPU struct {
	Call               []Register
	Syscall            []Register
	ExternCall         []Register
	ExternCallVolatile []Register
	Return             []Register
}
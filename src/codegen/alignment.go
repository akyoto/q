package codegen

const (
	alignFunction = 0x20
)

// alignment determines the machine code alignment of the entry point.
func (f *Function) alignment() uint8 {
	if f.FullName == "run.crash" || f.FullName == "run.exit" {
		return 0
	}

	return alignFunction
}
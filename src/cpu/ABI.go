package cpu

// ABI is the Application Binary Interface which defines the registers used in function calls.
type ABI struct {
	In        []Register
	Out       []Register
	Clobbered []Register
	Preserved []Register
}
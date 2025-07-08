package macho

const (
	BaseAddress  = 0x1000000
	SizeCommands = Segment64Size*4 + DyldInfoCommandSize + MainSize + DylinkerCommandSize + len(LinkerString) + DylibCommandSize + len(LibSystemString)
	HeaderEnd    = HeaderSize + SizeCommands
)
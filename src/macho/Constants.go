package macho

const (
	BaseAddress  = 0x1000000
	NumCommands  = 8
	SizeCommands = Segment64Size*3 + Section64Size + DyldInfoCommandSize + MainSize + BuildVersionSize + DylinkerCommandSize + len(LinkerString) + DylibCommandSize + len(LibSystemString)
	HeaderEnd    = HeaderSize + SizeCommands
)
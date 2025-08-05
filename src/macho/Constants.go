package macho

const (
	BaseAddress  = 0x1000000
	NumCommands  = 9
	SizeCommands = Segment64Size*4 +
		Section64Size +
		MainSize +
		BuildVersionSize +
		ChainedFixupsCommandSize +
		DylinkerCommandSize + len(LinkerString) +
		DylibCommandSize + len(LibSystemString)
	HeaderEnd = HeaderSize + SizeCommands
)
package macho

const (
	BaseAddress  = 0x1000000
	NumCommands  = 10
	NumSegments  = 4
	SizeCommands = NumSegments*Segment64Size +
		Section64Size +
		MainSize +
		BuildVersionSize +
		LinkeditDataCommandSize*2 +
		DylinkerCommandSize + len(LinkerString) +
		DylibCommandSize + len(LibSystemString)
	HeaderEnd = HeaderSize + SizeCommands
)
package macho

const (
	BaseAddress  = 0x1000000
	NumCommands  = 11
	NumSegments  = 4
	SizeCommands = NumSegments*Segment64Size +
		Section64Size +
		UuidSize +
		MainSize +
		BuildVersionSize +
		LinkeditDataCommandSize*2 +
		DylinkerCommandSize + len(LinkerString) +
		DylibCommandSize + len(LibSystemString)
	HeaderEnd = HeaderSize + SizeCommands
)
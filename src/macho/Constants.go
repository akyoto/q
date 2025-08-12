package macho

const (
	BaseAddress  = 0x100000000
	NumCommands  = 10
	NumSegments  = 3
	SizeCommands = NumSegments*Segment64Size +
		Section64Size +
		UuidSize +
		MainSize +
		BuildVersionSize +
		LinkeditDataCommandSize*2 +
		DylinkerCommandSize + len(LinkerString) +
		DylibCommandSize + len(LibSystemString)
	HashPageSize = 4096
	HeaderEnd    = HeaderSize + SizeCommands
)
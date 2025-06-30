package macho

const (
	BaseAddress  = 0x1000000
	SizeCommands = Segment64Size*3 + ThreadSize
	HeaderEnd    = HeaderSize + SizeCommands
)
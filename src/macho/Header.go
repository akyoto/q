package macho

const HeaderSize = 32

// Header contains general information.
type Header struct {
	Magic             uint32
	Architecture      CPU
	MicroArchitecture uint32
	Type              HeaderType
	NumCommands       uint32
	SizeCommands      uint32
	Flags             HeaderFlags
	Reserved          uint32
}
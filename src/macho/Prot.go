package macho

type Prot uint32

const (
	ProtReadable   Prot = 0x1
	ProtWritable   Prot = 0x2
	ProtExecutable Prot = 0x4
)
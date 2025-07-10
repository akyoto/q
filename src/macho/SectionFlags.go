package macho

type SectionFlags uint32

const (
	SectionPureInstructions SectionFlags = 0x80000000
)
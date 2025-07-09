package macho

type SectionFlags uint32

const (
	FlagPureInstructions SectionFlags = 0x80000000
)
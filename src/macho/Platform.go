package macho

type Platform uint32

const (
	PlatformInvalid Platform = iota
	PlatformMacOS
	PlatformIOS
)
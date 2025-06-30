package macho

const ThreadSize = 184

// Thread is a thread state load command.
type Thread struct {
	LoadCommand
	Len  uint32
	Type uint32
	Data [43]uint32
}
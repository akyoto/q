package exe

// Section represents some data within the executable.
type Section struct {
	Bytes        []byte
	Padding      int
	FileOffset   int
	MemoryOffset int
}
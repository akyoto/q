package exe

// Section represents some data within the executable that will also be loaded into memory.
type Section struct {
	Bytes        []byte
	FileOffset   int
	Padding      int
	MemoryOffset int
}
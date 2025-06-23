package exe

// Executable is a generic definition of the binary that later gets translated to OS-specific formats.
type Executable struct {
	Sections    []*Section
	headerEnd   int
	fileAlign   int
	memoryAlign int
}

// New creates a new executable.
func New(headerEnd int, fileAlign int, memoryAlign int) *Executable {
	return &Executable{
		headerEnd:   headerEnd,
		fileAlign:   fileAlign,
		memoryAlign: memoryAlign,
	}
}

// InitSections generates sections from raw byte slices.
func (exe *Executable) InitSections(raw ...[]byte) {
	exe.Sections = make([]*Section, len(raw))

	for i, data := range raw {
		exe.Sections[i] = &Section{Bytes: data}
	}

	exe.Update()
}

// Update recalculates all section offsets.
func (exe *Executable) Update() {
	first := exe.Sections[0]
	first.FileOffset, first.Padding = AlignPad(exe.headerEnd, exe.fileAlign)
	first.MemoryOffset = Align(exe.headerEnd, exe.memoryAlign)

	for i, section := range exe.Sections[1:] {
		previous := exe.Sections[i]
		section.FileOffset, section.Padding = AlignPad(previous.FileOffset+len(previous.Bytes), exe.fileAlign)
		section.MemoryOffset = Align(previous.MemoryOffset+len(previous.Bytes), exe.memoryAlign)
	}
}
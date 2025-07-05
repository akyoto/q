package exe

// Executable is a generic definition of the binary that later gets translated to OS-specific formats.
type Executable struct {
	Sections    []*Section
	headerEnd   int
	fileAlign   int
	memoryAlign int
	congruent   bool
}

// New creates a new executable.
func New(headerEnd int, fileAlign int, memoryAlign int, congruent bool, raw ...[]byte) *Executable {
	exe := &Executable{
		Sections:    make([]*Section, len(raw)),
		headerEnd:   headerEnd,
		fileAlign:   fileAlign,
		memoryAlign: memoryAlign,
		congruent:   congruent,
	}

	for i, section := range raw {
		if len(section) == 0 {
			section = []byte{0}
		}

		exe.Sections[i] = &Section{Bytes: section}
	}

	exe.Update()
	return exe
}

// Update recalculates all section offsets.
func (exe *Executable) Update() {
	first := exe.Sections[0]
	first.FileOffset, first.Padding = AlignPad(exe.headerEnd, exe.fileAlign)
	first.MemoryOffset = Align(exe.headerEnd, exe.memoryAlign)

	if exe.congruent && exe.fileAlign != exe.memoryAlign {
		first.MemoryOffset += first.FileOffset % exe.memoryAlign
	}

	for i, section := range exe.Sections[1:] {
		previous := exe.Sections[i]
		section.FileOffset, section.Padding = AlignPad(previous.FileOffset+len(previous.Bytes), exe.fileAlign)
		section.MemoryOffset = Align(previous.MemoryOffset+len(previous.Bytes), exe.memoryAlign)

		if exe.congruent && exe.fileAlign != exe.memoryAlign {
			section.MemoryOffset += section.FileOffset % exe.memoryAlign
		}
	}
}
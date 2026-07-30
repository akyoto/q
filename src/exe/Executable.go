package exe

// Executable is a generic definition of the binary that later gets translated to OS-specific formats.
type Executable struct {
	Sections     []*Section
	headerEnd    int
	fileAlign    int
	memoryAlign  int
	congruent    bool
	embedHeaders bool
}

// New creates a new executable.
func New(headerEnd int, fileAlign int, memoryAlign int, congruent bool, embedHeaders bool) *Executable {
	return &Executable{
		headerEnd:    headerEnd,
		fileAlign:    fileAlign,
		memoryAlign:  memoryAlign,
		congruent:    congruent,
		embedHeaders: embedHeaders,
	}
}

// AddSections adds the given byte slices as sections to the executable.
func (exe *Executable) AddSections(raw ...[]byte) {
	if exe.Sections == nil {
		exe.Sections = make([]*Section, len(raw))
	}

	for i, data := range raw {
		if len(data) == 0 {
			data = []byte{0}
		}

		exe.Sections[i] = &Section{Bytes: data}
	}

	exe.update()
}

// update recalculates all section offsets.
func (exe *Executable) update() {
	first := exe.Sections[0]

	if exe.embedHeaders {
		first.FileOffset, first.Padding = AlignPad(exe.headerEnd, 0x40)
		first.MemoryOffset = first.FileOffset
	} else {
		first.FileOffset, first.Padding = AlignPad(exe.headerEnd, exe.fileAlign)
		first.MemoryOffset = Align(exe.headerEnd, exe.memoryAlign)
	}

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
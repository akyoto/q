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
	for _, data := range raw {
		if len(data) == 0 {
			data = []byte{0}
		}

		index := len(exe.Sections)
		section := &Section{Bytes: data}
		exe.update(section, index)
		exe.Sections = append(exe.Sections, section)
	}
}

// update recalculates the section offsets.
func (exe *Executable) update(section *Section, index int) {
	if index == 0 {
		exe.updateFirst(section)
		return
	}

	previous := exe.Sections[index-1]
	section.FileOffset, section.Padding = AlignPad(previous.FileOffset+len(previous.Bytes), exe.fileAlign)
	section.MemoryOffset = Align(previous.MemoryOffset+len(previous.Bytes), exe.memoryAlign)

	if exe.congruent && exe.fileAlign != exe.memoryAlign {
		section.MemoryOffset += section.FileOffset % exe.memoryAlign
	}
}

// updateFirst recalculates the section offsets for the first section.
func (exe *Executable) updateFirst(section *Section) {
	if exe.embedHeaders {
		section.FileOffset, section.Padding = AlignPad(exe.headerEnd, 0x40)
		section.MemoryOffset = section.FileOffset
	} else {
		section.FileOffset, section.Padding = AlignPad(exe.headerEnd, exe.fileAlign)
		section.MemoryOffset = Align(exe.headerEnd, exe.memoryAlign)
	}

	if exe.congruent && exe.fileAlign != exe.memoryAlign {
		section.MemoryOffset += section.FileOffset % exe.memoryAlign
	}
}
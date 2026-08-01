package macho

import (
	"bytes"
	"encoding/binary"
	"strings"

	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/exe"
)

const (
	// Import formats
	DYLD_CHAINED_IMPORT          = 1
	DYLD_CHAINED_IMPORT_ADDEND   = 2
	DYLD_CHAINED_IMPORT_ADDEND64 = 3

	// Pointer formats
	DYLD_CHAINED_PTR_ARM64E              = 1
	DYLD_CHAINED_PTR_64                  = 2
	DYLD_CHAINED_PTR_32                  = 3
	DYLD_CHAINED_PTR_32_CACHE            = 4
	DYLD_CHAINED_PTR_32_FIRMWARE         = 5
	DYLD_CHAINED_PTR_64_OFFSET           = 6
	DYLD_CHAINED_PTR_ARM64E_OFFSET       = 7
	DYLD_CHAINED_PTR_ARM64E_KERNEL       = 7
	DYLD_CHAINED_PTR_64_KERNEL_CACHE     = 8
	DYLD_CHAINED_PTR_ARM64E_USERLAND     = 9
	DYLD_CHAINED_PTR_ARM64E_FIRMWARE     = 10
	DYLD_CHAINED_PTR_X86_64_KERNEL_CACHE = 11
	DYLD_CHAINED_PTR_ARM64E_USERLAND24   = 12
	DYLD_CHAINED_PTR_ARM64E_SHARED_CACHE = 13

	// Special start values
	DYLD_CHAINED_PTR_START_NONE  = 0xFFFF
	DYLD_CHAINED_PTR_START_LAST  = 0x8000
	DYLD_CHAINED_PTR_START_MULTI = 0x8000
)

// LinkEdit represents read-only data for the dynamic linker.
type LinkEdit struct {
	ChainedFixupsHeader
	ChainedStartsInImage
	ChainedStartsInSegment
}

// createLinkEditSegment creates the contents of the __LINKEDIT segment.
func createLinkEditSegment(libs dll.List, data *exe.Section, pageSize int) []byte {
	data.Bytes = exe.PadSlice(data.Bytes, 16)
	numFunctions := libs.CountFunctions()
	linkEditSize := ChainedFixupsHeaderSize + ChainedStartsInImageSize + ChainedStartsInSegmentSize
	importsOffset := exe.Align(linkEditSize, 16)
	importsSize := numFunctions * 8
	symbolsOffset := exe.Align(importsOffset+importsSize, 16)

	linkEdit := LinkEdit{
		ChainedFixupsHeader: ChainedFixupsHeader{
			FixupsVersion: 0,
			StartsOffset:  ChainedFixupsHeaderSize,
			ImportsOffset: uint32(importsOffset),
			SymbolsOffset: uint32(symbolsOffset),
			ImportsCount:  uint32(numFunctions),
			ImportsFormat: DYLD_CHAINED_IMPORT,
			SymbolsFormat: 0,
		},
		ChainedStartsInImage: ChainedStartsInImage{
			SegCount:      NumSegments,
			SegInfoOffset: [NumSegments]uint32{0, 0, ChainedStartsInImageSize, 0},
		},
		ChainedStartsInSegment: ChainedStartsInSegment{
			Size:            ChainedStartsInSegmentSize,
			PageSize:        uint16(pageSize),
			PointerFormat:   DYLD_CHAINED_PTR_64_OFFSET,
			SegmentOffset:   uint64(data.MemoryOffset),
			MaxValidPointer: 0,
			PageCount:       1,
			PageStarts:      [1]uint16{uint16(len(data.Bytes))},
		},
	}

	buffer := bytes.Buffer{}
	binary.Write(&buffer, binary.LittleEndian, &linkEdit.ChainedFixupsHeader)
	binary.Write(&buffer, binary.LittleEndian, &linkEdit.ChainedStartsInImage)
	binary.Write(&buffer, binary.LittleEndian, &linkEdit.ChainedStartsInSegment)

	// Imports
	exe.PadBuffer(&buffer, 16)
	symbols := strings.Builder{}
	libOrdinal := 1

	for lib := range libs.All() {
		for _, name := range lib.Functions {
			nameOffset := symbols.Len()
			symbols.WriteString(name)
			symbols.WriteByte(0)
			weakImport := 0
			dyldChainedImport := (nameOffset << 9) + (weakImport << 8) + libOrdinal
			binary.Write(&buffer, binary.LittleEndian, dyldChainedImport)
			data.Bytes = binary.LittleEndian.AppendUint64(data.Bytes, 0)
		}

		libOrdinal++
	}

	// Symbols
	exe.PadBuffer(&buffer, 16)
	buffer.WriteString(symbols.String())

	// Make sure the code signature that follows is 16-byte aligned
	return exe.PadSlice(buffer.Bytes(), 16)
}
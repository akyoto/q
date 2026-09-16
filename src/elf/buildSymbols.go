package elf

import (
	"cmp"
	"encoding/binary"
	"slices"
)

const (
	SymbolSize = 24
	funcGlobal = 0x12
)

// buildSymbols builds the symbol table and its string table from the function code offsets.
func buildSymbols(virtualAddress int64, codeLength int, labels map[string]int) (symtab []byte, strtab []byte) {
	names := make([]string, 0, len(labels))
	offsets := make(map[string]int, len(labels))

	for name, offset := range labels {
		names = append(names, name)
		offsets[name] = offset
	}

	slices.SortFunc(names, func(a, b string) int {
		aOffset := offsets[a]
		bOffset := offsets[b]

		if aOffset != bOffset {
			return aOffset - bOffset
		}

		return cmp.Compare(a, b)
	})

	strtab = []byte{0}
	index := make(map[string]int32, len(names))

	for _, name := range names {
		index[name] = int32(len(strtab))
		strtab = append(strtab, name...)
		strtab = append(strtab, 0)
	}

	symtab = make([]byte, SymbolSize*(len(names)+1))

	for i, name := range names {
		base := SymbolSize * (i + 1)
		offset := offsets[name]
		size := codeLength - offset

		if i+1 < len(names) {
			size = offsets[names[i+1]] - offset
		}

		binary.LittleEndian.PutUint32(symtab[base:], uint32(index[name]))
		symtab[base+4] = funcGlobal
		binary.LittleEndian.PutUint16(symtab[base+6:], 1)
		binary.LittleEndian.PutUint64(symtab[base+8:], uint64(virtualAddress+int64(offset)))
		binary.LittleEndian.PutUint64(symtab[base+16:], uint64(size))
	}

	return symtab, strtab
}
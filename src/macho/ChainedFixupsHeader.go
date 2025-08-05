package macho

const ChainedFixupsHeaderSize = 32

// ChainedFixupsHeader is the root of the chained fixups data,
// containing offsets to the starts, imports, and symbols tables.
type ChainedFixupsHeader struct {
	FixupsVersion uint32
	StartsOffset  uint32
	ImportsOffset uint32
	SymbolsOffset uint32
	ImportsCount  uint32
	ImportsFormat uint32
	SymbolsFormat uint32
	_             uint32
}
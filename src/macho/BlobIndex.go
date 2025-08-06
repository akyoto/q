package macho

const BlobIndexSize = 8

type BlobIndex struct {
	Type   uint32
	Offset uint32
}
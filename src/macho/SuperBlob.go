package macho

const SuperBlobSize = 12

type SuperBlob struct {
	Magic  uint32
	Length uint32
	Count  uint32
}
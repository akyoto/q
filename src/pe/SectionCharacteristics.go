package pe

type SectionCharacteristics uint32

const (
	IMAGE_SCN_CNT_CODE               SectionCharacteristics = 0x00000020
	IMAGE_SCN_CNT_INITIALIZED_DATA   SectionCharacteristics = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA SectionCharacteristics = 0x00000080
	IMAGE_SCN_LNK_COMDAT             SectionCharacteristics = 0x00001000
	IMAGE_SCN_MEM_DISCARDABLE        SectionCharacteristics = 0x02000000
	IMAGE_SCN_MEM_EXECUTE            SectionCharacteristics = 0x20000000
	IMAGE_SCN_MEM_READ               SectionCharacteristics = 0x40000000
	IMAGE_SCN_MEM_WRITE              SectionCharacteristics = 0x80000000
)
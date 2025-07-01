package pe

type Characteristics uint16

const (
	IMAGE_FILE_RELOCS_STRIPPED         Characteristics = 0x0001
	IMAGE_FILE_EXECUTABLE_IMAGE        Characteristics = 0x0002
	IMAGE_FILE_LINE_NUMS_STRIPPED      Characteristics = 0x0004
	IMAGE_FILE_LOCAL_SYMS_STRIPPED     Characteristics = 0x0008
	IMAGE_FILE_AGGRESSIVE_WS_TRIM      Characteristics = 0x0010
	IMAGE_FILE_LARGE_ADDRESS_AWARE     Characteristics = 0x0020
	IMAGE_FILE_BYTES_REVERSED_LO       Characteristics = 0x0080
	IMAGE_FILE_32BIT_MACHINE           Characteristics = 0x0100
	IMAGE_FILE_DEBUG_STRIPPED          Characteristics = 0x0200
	IMAGE_FILE_REMOVABLE_RUN_FROM_SWAP Characteristics = 0x0400
	IMAGE_FILE_NET_RUN_FROM_SWAP       Characteristics = 0x0800
	IMAGE_FILE_SYSTEM                  Characteristics = 0x1000
	IMAGE_FILE_DLL                     Characteristics = 0x2000
	IMAGE_FILE_UP_SYSTEM_ONLY          Characteristics = 0x4000
	IMAGE_FILE_BYTES_REVERSED_HI       Characteristics = 0x8000
)
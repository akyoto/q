package macho

type HeaderType uint32

const (
	TypeObject  HeaderType = 0x1
	TypeExecute HeaderType = 0x2
	TypeCore    HeaderType = 0x4
	TypeDylib   HeaderType = 0x6
	TypeBundle  HeaderType = 0x8
	TypeDsym    HeaderType = 0xA
)
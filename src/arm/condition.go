package arm

type condition uint8

const (
	EQ condition = iota
	NE
	CS
	CC
	MI
	PL
	VS
	VC
	HI
	LS
	GE
	LT
	GT
	LE
	AL
	NV
)
package arm

type condition uint8

const (
	EQ condition = iota
	NE
	HS // CS
	LO // CC
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
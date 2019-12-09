package assembler

import "github.com/akyoto/asm"

type instruction interface {
	Exec(*asm.Assembler)
	Name() string
	SetName(string)
	String() string
	Size() byte
}

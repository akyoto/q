package assembler

import (
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

func (a *Assembler) Return() {
	lastInstr := a.lastInstruction()

	if lastInstr != nil {
		// Avoid double return
		if lastInstr.Name() == mnemonics.RET {
			return
		}

		// If the previous instruction was a call,
		// change it to a jump.
		// if lastInstr.Name() == CALL {
		// 	lastInstr.SetName(JMP)
		// 	return
		// }
	}

	a.do(mnemonics.RET)
}

func (a *Assembler) Syscall() {
	a.do(mnemonics.SYSCALL)
}

func (a *Assembler) Call(label string) {
	a.doJump(mnemonics.CALL, label)
}

func (a *Assembler) Jump(label string) {
	a.doJump(mnemonics.JMP, label)
}

func (a *Assembler) JumpIfEqual(label string) {
	a.doJump(mnemonics.JE, label)
}

func (a *Assembler) JumpIfNotEqual(label string) {
	a.doJump(mnemonics.JNE, label)
}

func (a *Assembler) JumpIfLess(label string) {
	a.doJump(mnemonics.JL, label)
}

func (a *Assembler) JumpIfLessOrEqual(label string) {
	a.doJump(mnemonics.JLE, label)
}

func (a *Assembler) JumpIfGreater(label string) {
	a.doJump(mnemonics.JG, label)
}

func (a *Assembler) JumpIfGreaterOrEqual(label string) {
	a.doJump(mnemonics.JGE, label)
}

func (a *Assembler) IncreaseRegister(destination *register.Register) {
	a.doRegister(mnemonics.INC, destination)
}

func (a *Assembler) DecreaseRegister(destination *register.Register) {
	a.doRegister(mnemonics.DEC, destination)
}

func (a *Assembler) PushRegister(destination *register.Register) {
	a.doRegister(mnemonics.PUSH, destination)
}

func (a *Assembler) PopRegister(destination *register.Register) {
	a.doRegister(mnemonics.POP, destination)
}

func (a *Assembler) DivRegister(destination *register.Register) {
	a.doRegister(mnemonics.DIV, destination)
}

func (a *Assembler) SignExtendToDX(destination *register.Register) {
	a.doRegister(mnemonics.CDQ, destination)
}

func (a *Assembler) MoveRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegisterRegister(mnemonics.MOV, destination, source)
}

func (a *Assembler) MoveRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(mnemonics.MOV, destination, number)
}

func (a *Assembler) StoreNumber(destination *register.Register, offset byte, byteCount byte, number uint64) {
	a.doMemoryNumber(mnemonics.STORE, destination, offset, byteCount, number)
}

func (a *Assembler) StoreRegister(destination *register.Register, offset byte, byteCount byte, source *register.Register) {
	a.doMemoryRegister(mnemonics.STORE, destination, offset, byteCount, source)
}

func (a *Assembler) LoadRegister(destination *register.Register, source *register.Register, offset byte, byteCount byte) {
	a.doRegisterMemory(mnemonics.LOAD, destination, source, offset, byteCount)
}

func (a *Assembler) MoveRegisterAddress(destination *register.Register, address uint32) {
	a.doRegisterAddress(mnemonics.MOV, destination, address)
}

func (a *Assembler) CompareRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegisterRegister(mnemonics.CMP, destination, source)
}

func (a *Assembler) CompareRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(mnemonics.CMP, destination, number)
}

func (a *Assembler) AddRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegisterRegister(mnemonics.ADD, destination, source)
}

func (a *Assembler) AddRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(mnemonics.ADD, destination, number)
}

func (a *Assembler) SubRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegisterRegister(mnemonics.SUB, destination, source)
}

func (a *Assembler) SubRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(mnemonics.SUB, destination, number)
}

func (a *Assembler) MulRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegisterRegister(mnemonics.MUL, destination, source)
}

func (a *Assembler) MulRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(mnemonics.MUL, destination, number)
}

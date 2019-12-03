package assembler

import "github.com/akyoto/q/build/register"

func (a *Assembler) Return() {
	lastInstr := a.lastInstruction()

	if lastInstr != nil && lastInstr.Name() == RET {
		return
	}

	a.do(RET)
}

func (a *Assembler) Syscall() {
	a.do(SYSCALL)
}

func (a *Assembler) Call(label string) {
	a.doLabel(CALL, label)
}

func (a *Assembler) Jump(label string) {
	a.doLabel(JMP, label)
}

func (a *Assembler) JumpIfEqual(label string) {
	a.doLabel(JE, label)
}

func (a *Assembler) JumpIfNotEqual(label string) {
	a.doLabel(JNE, label)
}

func (a *Assembler) JumpIfLess(label string) {
	a.doLabel(JL, label)
}

func (a *Assembler) JumpIfLessOrEqual(label string) {
	a.doLabel(JLE, label)
}

func (a *Assembler) JumpIfGreater(label string) {
	a.doLabel(JG, label)
}

func (a *Assembler) JumpIfGreaterOrEqual(label string) {
	a.doLabel(JGE, label)
}

func (a *Assembler) IncreaseRegister(destination *register.Register) {
	a.doRegister1(INC, destination)
}

func (a *Assembler) DecreaseRegister(destination *register.Register) {
	a.doRegister1(DEC, destination)
}

func (a *Assembler) PushRegister(destination *register.Register) {
	lastInstr := a.lastInstruction()

	// If the last instruction popped the same register,
	// it is nearly equivalent to not popping the register in the first place.
	if lastInstr != nil && lastInstr.Name() == POP && lastInstr.(*register1).Destination == destination {
		a.removeLastInstruction()
		return
	}

	a.doRegister1(PUSH, destination)
}

func (a *Assembler) PopRegister(destination *register.Register) {
	a.doRegister1(POP, destination)
}

func (a *Assembler) MoveRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegister2(MOV, destination, source)
}

func (a *Assembler) MoveRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(MOV, destination, number)
}

func (a *Assembler) MoveRegisterAddress(destination *register.Register, address uint32) {
	a.doRegisterAddress(MOV, destination, address)
}

func (a *Assembler) CompareRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegister2(CMP, destination, source)
}

func (a *Assembler) CompareRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(CMP, destination, number)
}

func (a *Assembler) AddRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegister2(ADD, destination, source)
}

func (a *Assembler) AddRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(ADD, destination, number)
}

func (a *Assembler) SubRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegister2(SUB, destination, source)
}

func (a *Assembler) SubRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(SUB, destination, number)
}

func (a *Assembler) MulRegisterRegister(destination *register.Register, source *register.Register) {
	a.doRegister2(MUL, destination, source)
}

func (a *Assembler) MulRegisterNumber(destination *register.Register, number uint64) {
	a.doRegisterNumber(MUL, destination, number)
}

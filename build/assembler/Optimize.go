package assembler

import (
	"github.com/akyoto/q/build/assembler/instructions"
	"github.com/akyoto/q/build/assembler/mnemonics"
)

// Optimize optimizes the assembler instructions for improved performance.
func (a *Assembler) Optimize() {
	for index, instr := range a.Instructions {
		// --------------------------------------------
		// Optimize a pop followed by a move
		// to a single pop.
		// --------------------------------------------
		// pop reg1
		// mov reg2, reg1
		// --------------------------------------------
		// pop reg2
		// --------------------------------------------
		// This optimization is only correct
		// if reg1 is not used in the remaining code.
		// --------------------------------------------
		if instr.Name() == mnemonics.POP {
			pop := instr.(*instructions.Register)
			nextInstr, ok := a.Instructions[index+1].(*instructions.RegisterRegister)

			if !ok {
				continue
			}

			if nextInstr.Mnemonic != mnemonics.MOV {
				continue
			}

			source := nextInstr.Source

			if source != pop.Destination {
				continue
			}

			remainingCode := a.Instructions[index+2:]
			canOptimize := true

			for _, checkInstr := range remainingCode {
				reg1, ok := checkInstr.(*instructions.Register)

				if ok {
					if reg1.Destination == source && reg1.Mnemonic != mnemonics.POP {
						canOptimize = false
						break
					}

					continue
				}

				reg2, ok := checkInstr.(*instructions.RegisterRegister)

				if ok {
					if reg2.Source == source {
						canOptimize = false
						break
					}

					if reg2.Destination == source {
						// Overwriting the register means it's safe to optimize
						if reg2.Mnemonic == mnemonics.MOV {
							break
						}

						canOptimize = false
						break
					}

					continue
				}

				regNumber, ok := checkInstr.(*instructions.RegisterNumber)

				if ok {
					if regNumber.Destination == source && regNumber.Mnemonic != mnemonics.MOV {
						canOptimize = false
						break
					}

					continue
				}
			}

			if !canOptimize {
				continue
			}

			pop.Destination = nextInstr.Destination
			a.Instructions = append(a.Instructions[:index+1], remainingCode...)
		}
	}
}

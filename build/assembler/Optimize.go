package assembler

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
		if instr.Name() == POP {
			pop := instr.(*register1)
			nextInstr, ok := a.Instructions[index+1].(*register2)

			if !ok {
				continue
			}

			if nextInstr.Mnemonic != MOV {
				continue
			}

			source := nextInstr.Source

			if source != pop.Destination {
				continue
			}

			remainingCode := a.Instructions[index+2:]
			canOptimize := true

			for _, checkInstr := range remainingCode {
				reg1, ok := checkInstr.(*register1)

				if ok {
					if reg1.Destination == source && reg1.Mnemonic != POP {
						canOptimize = false
						break
					}

					continue
				}

				reg2, ok := checkInstr.(*register2)

				if ok {
					if reg2.Source == source || (reg2.Destination == source && reg2.Mnemonic != MOV) {
						canOptimize = false
						break
					}

					continue
				}

				regNumber, ok := checkInstr.(*registerNumber)

				if ok {
					if regNumber.Destination == source && regNumber.Mnemonic != MOV {
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

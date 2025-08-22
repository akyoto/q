package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// needsRegister returns true if the value requires a register.
func (f *Function) needsRegister(s *Step) bool {
	typ := s.Value.Type()

	if typ == types.Void {
		return false
	}

	_, isStruct := typ.(*types.Struct)

	if isStruct {
		return false
	}

	_, isTuple := typ.(*types.Tuple)

	if isTuple {
		return false
	}

	users := s.Value.Users()

	if len(users) == 0 {
		return false
	}

	switch instr := s.Value.(type) {
	case *ssa.BinaryOp:
		return !instr.Op.IsComparison()
	case *ssa.Int:
		if len(users) == 1 {
			// Check if we can encode single-use integers as immediates
			// directly embedded in the instruction itself rather than
			// requiring an extra register and a move.
			return !f.canEncodeNumber(users[0], instr)
		}
	}

	return true
}
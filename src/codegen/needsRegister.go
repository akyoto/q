package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// needsRegister returns true if the value requires a register.
func (f *Function) needsRegister(s *Step) bool {
	typ := types.Unwrap(s.Value.Type())

	if typ == types.Void {
		return false
	}

	_, isPhi := s.Value.(*ssa.Phi)

	if isPhi {
		return true
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
		if instr.Op.IsComparison() {
			next := f.Steps[s.Index+1]
			branch, isBranch := next.Value.(*ssa.Branch)
			return !isBranch || !slices.Contains(branch.Inputs(), s.Value)
		}

		return true
	case *ssa.Cas:
		return false
	case *ssa.Int:
		// Check if we can encode zero as an immediate directly
		// embedded in the instructions rather than requiring
		// an extra register and a move.
		if f.arch.sharedImmediate(instr) {
			for _, user := range users {
				if !f.arch.canEncodeNumber(user, instr) {
					return true
				}
			}

			return false
		}

		if len(users) == 1 {
			return !f.arch.canEncodeNumber(users[0], instr)
		}
	case *ssa.Memory:
		return false
	}

	return true
}
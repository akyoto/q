package optimizer

import (
	"slices"

	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
)

// Reorder moves values closer to their first use.
func Reorder(ir ssa.IR) {
	for _, block := range ir.Blocks {
		for start, instr := range slices.Backward(block.Instructions) {
			switch instr.(type) {
			case *ssa.Bytes, *ssa.Data, *ssa.Int:
			default:
				continue
			}

			users := instr.Users()

			if len(users) == 0 {
				continue
			}

			firstUser := users[0]
			end := block.Index(firstUser)

			if end <= start+1 {
				continue
			}

			for end > start && slices.Contains(block.Instructions[end-1].Users(), firstUser) {
				end--
			}

			set.BringToBack(block.Instructions[start:end], 0)
		}
	}
}
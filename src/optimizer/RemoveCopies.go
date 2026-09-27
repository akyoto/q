package optimizer

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// RemoveCopies replaces copies with their actual values.
func RemoveCopies(ir *ssa.IR) {
	for _, block := range ir.Blocks {
		for i, value := range block.Instructions {
			copy, isCopy := value.(*ssa.Copy)

			if !isCopy {
				continue
			}

			typ := copy.Value.Type()

			if copy.Typ != typ {
				continue
			}

			_, isResource := typ.(*types.Resource)

			if isResource {
				continue
			}

			ir.ReplaceAll(copy, copy.Value)
			block.RemoveAt(i)
		}

		block.RemoveNilValues()
	}
}
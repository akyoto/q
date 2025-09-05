package core

import "git.urbach.dev/cli/q/src/ssa"

// removeCopies removes copy operations.
func (f *Function) removeCopies() {
	var replacements map[ssa.Value]ssa.Value

	for _, block := range f.Blocks {
		if block.Loop != nil {
			continue
		}

		for i, value := range block.Instructions {
			copy, isCopy := value.(*ssa.Copy)

			if !isCopy {
				continue
			}

			if replacements == nil {
				replacements = make(map[ssa.Value]ssa.Value)
			}

			original := copy.Value

			for {
				originalCopy, originalIsCopy := original.(*ssa.Copy)

				if !originalIsCopy {
					break
				}

				original = originalCopy.Value
			}

			replacements[copy] = original
			block.RemoveAt(i)
		}

		block.RemoveNilValues()
	}

	for old, new := range replacements {
		f.IR.ReplaceAll(old, new)
	}
}
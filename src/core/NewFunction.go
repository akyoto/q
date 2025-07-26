package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/ssa"
)

// NewFunction creates a new function.
func NewFunction(name string, pkg string, file *fs.File) *Function {
	fullName := fmt.Sprintf("%s.%s", pkg, name)

	return &Function{
		Name:    name,
		Package: pkg,
		File:    file,
		IR: ssa.IR{
			Blocks: []*ssa.Block{
				{
					Label:        fullName,
					Instructions: make([]ssa.Value, 0, 8),
				},
			},
		},
		Function: codegen.Function{
			FullName: fullName,
			Assembler: asm.Assembler{
				Instructions: make([]asm.Instruction, 0, 8),
			},
		},
	}
}
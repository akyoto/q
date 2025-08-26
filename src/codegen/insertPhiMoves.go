package codegen

import "git.urbach.dev/cli/q/src/asm"

// insertPhiMoves moves all live values that are part of a Phi instruction
// from their current register to the Phi target register.
// It must be called right before a Jump instruction.
func (f *Function) insertPhiMoves(step *Step) {
	for _, live := range step.Live {
		if live.Phi != nil && live.Register != live.Phi.Register {
			f.Assembler.Append(&asm.Move{
				Destination: live.Phi.Register,
				Source:      live.Register,
			})
		}
	}
}
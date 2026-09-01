package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/cpu"
)

// phiMoveClobbered returns the destinations of phi moves.
func phiMoveClobbered(step *Step) []cpu.Register {
	var clobbered []cpu.Register

	for _, live := range step.Live {
		for phi := range live.Phis.All() {
			if live.Register == phi.Register {
				continue
			}

			if slices.Contains(clobbered, phi.Register) {
				continue
			}

			if !slices.Contains(phi.Block.Predecessors, step.Block) {
				continue
			}

			clobbered = append(clobbered, phi.Register)
		}
	}

	return clobbered
}
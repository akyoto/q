package codegen

import (
	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/cpu"
)

// ScheduleMoves orders the given moves using topological sort and resolves cycles with a temporary register.
func ScheduleMoves(moves []*asm.Move, free []cpu.Register) ([]*asm.Move, bool) {
	remaining := make([]*asm.Move, 0, len(moves))

	for _, move := range moves {
		if move.Source != move.Destination {
			remaining = append(remaining, move)
		}
	}

	result := make([]*asm.Move, 0, len(moves))

start:
	for len(remaining) > 0 {
		sources := bitSet(0)

		for _, move := range remaining {
			sources.Set(move.Source)
		}

		for i, move := range remaining {
			if sources.Has(move.Destination) {
				continue
			}

			result = append(result, move)
			remaining = append(remaining[:i], remaining[i+1:]...)
			continue start
		}

		// Now there are only cycles remaining.
		visited := make([]bool, len(remaining))
		current := 0

		for !visited[current] {
			visited[current] = true
			next := -1

			for j, move := range remaining {
				if move.Source == remaining[current].Destination {
					next = j
					break
				}
			}

			current = next
		}

		// Use a temporary register to break the cycle.
		move := remaining[current]
		used := bitSet(0)

		for _, move := range remaining {
			used.Set(move.Source)
			used.Set(move.Destination)
		}

		temp := cpu.Register(-1)

		for _, reg := range free {
			if !used.Has(reg) {
				temp = reg
				break
			}
		}

		if temp == -1 {
			return nil, false
		}

		result = append(result, &asm.Move{Destination: temp, Source: move.Source})
		remaining = append(remaining[:current], remaining[current+1:]...)
		remaining = append(remaining, &asm.Move{Destination: move.Destination, Source: temp})
	}

	return result, true
}
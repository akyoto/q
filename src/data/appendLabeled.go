package data

import (
	"slices"

	"git.urbach.dev/cli/q/src/exe"
)

// appendLabeled adds labeled data, sorted by label, each padded to its size's alignment.
func appendLabeled(final []byte, positions map[string]int, labeled map[string][]byte) []byte {
	keys := make([]string, 0, len(labeled))

	for key := range labeled {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		content := labeled[key]
		final = exe.PadSlice(final, len(content))
		positions[key] = len(final)
		final = append(final, content...)
	}

	return final
}
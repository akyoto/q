package data

import (
	"slices"

	"git.urbach.dev/cli/q/src/exe"
)

// appendMutable adds data that is mutable and not subject to string interning.
func (data *Data) appendMutable(final []byte, positions map[string]int) []byte {
	keys := make([]string, 0, len(data.Mutable))

	for key := range data.Mutable {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	for _, key := range keys {
		content := data.Mutable[key]
		final = exe.PadSlice(final, len(content))
		positions[key] = len(final)
		final = append(final, content...)
	}

	return final
}
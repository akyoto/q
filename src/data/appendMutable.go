package data

import (
	"slices"
	"sort"

	"git.urbach.dev/cli/q/src/exe"
)

func (data *Data) appendMutable(final []byte, positions map[string]int) []byte {
	keys := make([]string, 0, len(data.Mutable))

	for key := range data.Mutable {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		a := data.Mutable[keys[i]]
		b := data.Mutable[keys[j]]

		if len(a) != len(b) {
			return len(a) > len(b)
		}

		return slices.Compare(a, b) == -1
	})

	for _, key := range keys {
		content := data.Mutable[key]
		_, padding := exe.AlignPad(len(final), len(content))

		for range padding {
			final = append(final, 0)
		}

		positions[key] = len(final)
		final = append(final, content...)
	}

	return final
}
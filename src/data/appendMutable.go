package data

import (
	"slices"
	"sort"
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
		positions[key] = len(final)
		final = append(final, data.Mutable[key]...)
	}

	return final
}
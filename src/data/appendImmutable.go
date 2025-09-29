package data

import (
	"bytes"
	"slices"
	"sort"
)

func (data *Data) appendImmutable(final []byte, positions map[string]int) []byte {
	keys := make([]string, 0, len(data.Immutable))

	for key := range data.Immutable {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		a := data.Immutable[keys[i]]
		b := data.Immutable[keys[j]]

		if len(a) != len(b) {
			return len(a) > len(b)
		}

		return slices.Compare(a, b) == -1
	})

	for _, key := range keys {
		raw := data.Immutable[key]
		position := bytes.Index(final, raw)

		if position != -1 {
			positions[key] = position
		} else {
			positions[key] = len(final)
			final = append(final, raw...)
		}
	}

	return final
}
package data

import (
	"bytes"
	"sort"
)

// Finalize returns the final raw data slice and a map of labels with their respective indices.
// It will try to reuse existing data whenever possible.
func (data Data) Finalize() ([]byte, map[string]int) {
	var (
		keys      = make([]string, 0, len(data))
		positions = make(map[string]int, len(data))
		capacity  = 0
	)

	for key, value := range data {
		keys = append(keys, key)
		capacity += len(value)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return len(data[keys[i]]) > len(data[keys[j]])
	})

	final := make([]byte, 0, capacity)

	for _, key := range keys {
		raw := data[key]
		position := bytes.Index(final, raw)

		if position != -1 {
			positions[key] = position
		} else {
			positions[key] = len(final)
			final = append(final, raw...)
		}
	}

	return final, positions
}
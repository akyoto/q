package data

import (
	"slices"

	"git.urbach.dev/cli/q/src/exe"
)

// TLSLayout assigns each label an offset within a contiguous thread-local block.
func TLSLayout(sizes map[string]int) (map[string]int, int) {
	offsets := make(map[string]int, len(sizes))
	labels := make([]string, 0, len(sizes))
	offset := 0

	for label := range sizes {
		labels = append(labels, label)
	}

	slices.Sort(labels)

	for _, label := range labels {
		offset += exe.Pad(offset, sizes[label])
		offsets[label] = offset
		offset += sizes[label]
	}

	return offsets, offset
}
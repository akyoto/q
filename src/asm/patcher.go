package asm

import "slices"

// patcher maintains patches to a collection of bytes.
type patcher struct {
	code         []byte
	labels       map[string]int
	earlyPatches []*patch
	latePatches  []*patch
}

// ApplyPatches applies the given patches to the code.
func (p *patcher) ApplyPatches(patches []*patch) {
restart:
	for i, patch := range patches {
		start := patch.start
		end := patch.end
		oldSize := end - start
		new := patch.apply(p.code[start:end])
		newSize := len(new)
		shift := newSize - oldSize

		if shift == 0 {
			continue
		}

		patch.end += shift

		for label, address := range p.labels {
			if address >= end {
				p.labels[label] += shift
			}
		}

		for _, following := range p.earlyPatches[i+1:] {
			if following.start >= end {
				following.start += shift
				following.end += shift
			}
		}

		for _, following := range p.latePatches[i+1:] {
			if following.start >= end {
				following.start += shift
				following.end += shift
			}
		}

		left := p.code[:start]
		right := p.code[end:]
		p.code = slices.Concat(left, new, right)
		goto restart
	}
}

// PatchLast4Bytes creates a late patch for the last 4 bytes.
func (p *patcher) PatchLast4Bytes() *patch {
	patch := &patch{
		start: len(p.code) - 4,
		end:   len(p.code),
	}

	p.latePatches = append(p.latePatches, patch)
	return patch
}
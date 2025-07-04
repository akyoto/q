package build

// cacheLineSize is the smallest unit of data that can be transferred between the RAM and the CPU cache.
const cacheLineSize = 64

// FileAlign returns the file alignment.
func (build *Build) FileAlign() int {
	if build.OS == Linux {
		return cacheLineSize
	}

	return build.MemoryAlign()
}
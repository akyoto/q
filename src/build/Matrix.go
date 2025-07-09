package build

// Matrix calls the given function with every possible combination of operating systems and architectures.
func (b *Build) Matrix(call func(*Build)) {
	systems := []OS{Linux, Mac, Windows}
	architectures := []Arch{ARM, X86}

	for _, os := range systems {
		b.OS = os

		for _, arch := range architectures {
			b.Arch = arch
			call(b)
		}
	}
}
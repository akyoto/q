package config

// Matrix calls the given function with every possible combination of operating systems and architectures.
func (build *Build) Matrix(call func(*Build)) {
	systems := []OS{Linux, Mac, Windows}
	architectures := []Arch{ARM, X86}
	tmp := *build

	for _, os := range systems {
		tmp.OS = os

		for _, arch := range architectures {
			tmp.Arch = arch
			call(&tmp)
		}
	}
}
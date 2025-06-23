package ssa

type Arguments struct {
	Args []Value
}

func (v *Arguments) Dependencies() []Value {
	return v.Args
}

func (a Arguments) Equals(b Arguments) bool {
	if len(a.Args) != len(b.Args) {
		return false
	}

	for i := range a.Args {
		if !a.Args[i].Equals(b.Args[i]) {
			return false
		}
	}

	return true
}
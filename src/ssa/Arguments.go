package ssa

type Arguments []Value

func (v Arguments) Dependencies() []Value {
	return v
}

func (a Arguments) Equals(b Arguments) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}
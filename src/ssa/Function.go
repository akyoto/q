package ssa

type Function struct {
	UniqueName string
	Liveness
	HasToken
}

func (v *Function) Dependencies() []Value {
	return nil
}

func (a *Function) Equals(v Value) bool {
	b, sameType := v.(*Function)

	if !sameType {
		return false
	}

	return a.UniqueName == b.UniqueName
}

func (v *Function) IsConst() bool {
	return true
}

func (v *Function) String() string {
	return v.UniqueName
}
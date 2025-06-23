package ssa

import "fmt"

type Return struct {
	Arguments
	HasToken
}

func (a *Return) AddUse(user Value) { panic("return is not a value") }
func (a *Return) Alive() int        { return 0 }

func (a *Return) Equals(v Value) bool {
	b, sameType := v.(*Return)

	if !sameType {
		return false
	}

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

func (v *Return) IsConst() bool {
	return false
}

func (v *Return) String() string {
	return fmt.Sprintf("return %v", v.Args)
}
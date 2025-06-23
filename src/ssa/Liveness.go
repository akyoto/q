package ssa

type Liveness struct {
	alive int
}

func (v *Liveness) AddUse(user Value) {
	v.alive++
}

func (v *Liveness) Alive() int {
	return v.alive
}
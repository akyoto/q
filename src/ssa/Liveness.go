package ssa

type Liveness struct {
	users []Value
}

func (v *Liveness) AddUser(user Value) {
	v.users = append(v.users, user)
}

func (v *Liveness) Users() []Value {
	return v.users
}
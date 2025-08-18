package ssa

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

func (v *Liveness) AddUser(user Value) {
	v.users = append(v.users, user)
}

func (v *Liveness) RemoveUser(user Value) {
	for i, search := range v.users {
		if search == user {
			v.users = append(v.users[:i], v.users[i+1:]...)
			return
		}
	}
}

func (v *Liveness) Users() []Value {
	return v.users
}
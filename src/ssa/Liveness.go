package ssa

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

func (v *Liveness) addUser(user Value) { v.users = append(v.users, user) }
func (v *Liveness) Users() []Value     { return v.users }
package ssa

type HasLiveness interface {
	AddUser(Value)
	ReplaceUser(Value, Value)
	Users() []Value
}

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

func (v *Liveness) AddUser(user Value) { v.users = append(v.users, user) }
func (v *Liveness) ReplaceUser(old Value, new Value) {
	for i, user := range v.users {
		if user == old {
			v.users[i] = new
		}
	}
}
func (v *Liveness) Users() []Value { return v.users }
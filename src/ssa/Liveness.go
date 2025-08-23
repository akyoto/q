package ssa

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

// AddUser adds a new user of the value.
func (l *Liveness) AddUser(user Value) {
	l.users = append(l.users, user)
}

// RemoveUser removes a user of the value.
func (l *Liveness) RemoveUser(user Value) {
	for i, search := range l.users {
		if search == user {
			l.users = append(l.users[:i], l.users[i+1:]...)
			return
		}
	}
}

// Users returns the users of the value.
func (l *Liveness) Users() []Value {
	return l.users
}
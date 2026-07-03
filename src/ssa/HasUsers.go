package ssa

type HasUsers interface {
	// AddUser adds a new user of this value.
	AddUser(Value)

	// RemoveUser removes an existing user of this value.
	RemoveUser(Value)

	// Users returns all values that reference this value as an input.
	Users() []Value
}
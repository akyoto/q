package ssa

type HasLiveness interface {
	AddUser(Value)
	Users() []Value
}
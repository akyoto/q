package ssa

type NoLiveness struct{}

func (a *NoLiveness) AddUser(user Value) { panic("value does not have liveness") }
func (a *NoLiveness) CountUsers() int    { return 0 }
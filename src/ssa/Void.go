package ssa

import "git.urbach.dev/cli/q/src/types"

// Void is the absence of a value.
type Void struct{}

func (v *Void) AddUser(Value)    { panic("can not be used as a dependency") }
func (v *Void) IsConst() bool    { return false }
func (v *Void) RemoveUser(Value) { panic("can not be used as a dependency") }
func (v *Void) Type() types.Type { return types.Void }
func (v *Void) Users() []Value   { return nil }
package core

import "git.urbach.dev/cli/q/src/ssa"

// functionState contains temporary compiler state.
type functionState struct {
	loopStack
	constantStack []*Constant
	valueToStruct map[ssa.Value]*ssa.Struct
}
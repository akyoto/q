package core

import (
	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/ssa"
)

// Function is the smallest unit of code.
type Function struct {
	ssa.IR
	codegen.Function
	functionDependencies
	functionIdentity
	functionOverloads
	functionParameters
	functionState
	Env *Environment
	Err error
}
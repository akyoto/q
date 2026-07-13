package core

import "git.urbach.dev/cli/q/src/ssa"

// needsImplicitReturn returns true if it needs to end with a return statement.
func (f *Function) needsImplicitReturn() bool {
	if f == f.Env.Init || f.FullName == "run.crash" || f.FullName == "run.exit" {
		return false
	}

	lastBlock := f.Block()

	if lastBlock.Loop != nil {
		return false
	}

	_, endsWithReturn := lastBlock.Last().(*ssa.Return)
	return !endsWithReturn
}
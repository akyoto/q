package core

import "git.urbach.dev/cli/q/src/ssa"

// needsReturn returns true if it needs to end with a return statement.
func (f *Function) needsReturn() bool {
	if f == f.Env.Init || f.FullName == "run.exit" {
		return false
	}

	lastBlock := f.Block()

	if lastBlock.Loop != nil {
		return false
	}

	_, endsWithReturn := lastBlock.Last().(*ssa.Return)
	return !endsWithReturn
}
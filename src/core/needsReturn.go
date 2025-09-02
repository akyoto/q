package core

import "git.urbach.dev/cli/q/src/ssa"

// needsReturn returns true if it needs to end with a return statement.
func (f *Function) needsReturn() bool {
	if f == f.Env.Init || f.FullName == "os.exit" {
		return false
	}

	_, endsWithReturn := f.Block().Last().(*ssa.Return)
	return !endsWithReturn
}
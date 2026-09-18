package core

import "git.urbach.dev/cli/q/src/ssa"

// hasLiveUser returns true if the value has a user that actually reads the value.
func (f *Function) hasLiveUser(value ssa.Value) bool {
	for _, user := range value.Users() {
		copy, isCopy := user.(*ssa.Copy)

		if !isCopy || copy.Read {
			return true
		}

		if !f.Env.Build.LintDeadCode {
			return true
		}
	}

	return false
}
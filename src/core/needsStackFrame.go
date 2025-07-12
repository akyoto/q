package core

import "git.urbach.dev/cli/q/src/config"

// needsStackFrame returns true if the function needs a stack frame.
func (f *Function) needsStackFrame() bool {
	return f.All.Build.Arch == config.ARM && !f.IsLeaf() && f != f.All.Init
}
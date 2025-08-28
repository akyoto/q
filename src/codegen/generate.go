package codegen

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// generate executes all steps.
func (f *Function) generate() {
	f.enter()

	for _, step := range f.Steps {
		f.execute(step)
	}

	if len(f.Steps) > 0 {
		_, lastIsReturn := f.Steps[len(f.Steps)-1].Value.(*ssa.Return)

		if lastIsReturn {
			return
		}
	}

	f.leave()
}
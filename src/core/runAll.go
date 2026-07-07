package core

import (
	"sort"

	"git.urbach.dev/cli/q/src/ssa"
)

// runAll will run the specified function from each package if it exists.
func (f *Function) runAll(funcName string) {
	keys := make([]string, 0, len(f.Env.Packages))

	for name := range f.Env.Packages {
		keys = append(keys, name)
	}

	sort.Strings(keys)

	for _, name := range keys {
		pkg := f.Env.Packages[name]

		if pkg.Name == "run" {
			continue
		}

		fn, hasFunc := pkg.Functions[funcName]

		if !hasFunc {
			continue
		}

		ssaFunc := &ssa.Function{
			FunctionRef: fn,
			Typ:         fn.Type,
		}

		f.call(ssaFunc, nil, ssa.Source{})
	}
}
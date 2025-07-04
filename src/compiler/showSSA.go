package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

// showSSA shows the SSA IR.
func showSSA(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		ansi.Yellow.Println(f.UniqueName + ":")

		for i, step := range f.Steps {
			ansi.Dim.Printf("%%%d = ", i)
			fmt.Print(step.Value.String())
			ansi.Dim.Printf(" %s %s %s\n", step.Value.Type().Name(), step.Register, step.Live)
		}

		fmt.Println()
	})
}
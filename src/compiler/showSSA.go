package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

// showSSA shows the SSA IR.
func showSSA(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		ansi.Yellow.Println(f.FullName + ":")

		for _, step := range f.Steps {
			_, isLabel := step.Value.(*codegen.Label)

			if isLabel {
				fmt.Println(step.Value.String() + ":")
				continue
			}

			ansi.Dim.Printf("%p = ", step.Value)
			fmt.Print(step.Value.String())
			ansi.Dim.Printf(" %s %s %s\n", step.Value.Type().Name(), step.Register, step.Live)
		}

		fmt.Println()
	})
}
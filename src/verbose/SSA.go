package verbose

import (
	_ "embed"
	"fmt"

	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

//go:embed SSA.txt
var HeaderSSA string

// SSA shows the SSA IR.
func SSA(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		if filter(f.FullName, f.Env.Build.Filter) {
			return
		}

		ansi.Yellow.Println(f.FullName + ":")

		for _, step := range f.Steps {
			_, isLabel := step.Value.(*codegen.Label)

			if isLabel {
				fmt.Print(step.Value.String() + ":")

				for _, pre := range step.Block.Predecessors {
					ansi.Dim.Print(" ⇠ ")
					ansi.Dim.Print(pre)
				}

				fmt.Println()
				continue
			}

			ansi.Dim.Printf("  %p = ", step.Value)
			fmt.Print(step.Value.String())
			ansi.Dim.Printf(" %s [%s] ", step.Value.Type().Name(), step.Register)

			for identifier := range step.Block.IdentifiersFor(step.Value) {
				ansi.Dim.Printf("%s ", identifier)
			}

			for _, live := range step.Live {
				ansi.Dim.Printf("%p ", live.Value)
			}

			fmt.Println()
		}

		fmt.Println()
	})
}
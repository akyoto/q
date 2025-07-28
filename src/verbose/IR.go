package verbose

import (
	_ "embed"
	"fmt"

	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

//go:embed IR.txt
var bannerIR string

// IR shows the SSA IR.
func IR(root *core.Function) {
	Title(bannerIR)

	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		ansi.Yellow.Println(f.FullName + ":")

		for _, step := range f.Steps {
			_, isLabel := step.Value.(*codegen.Label)

			if isLabel {
				fmt.Println(step.Value.String() + ":")
				continue
			}

			ansi.Dim.Printf("  %p = ", step.Value)
			fmt.Print(step.Value.String())
			ansi.Dim.Printf(" %s [%s] ", step.Value.Type().Name(), step.Register)

			for _, live := range step.Live {
				ansi.Dim.Printf("%p ", live.Value)
			}

			fmt.Println()
		}

		fmt.Println()
	})
}
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

		for _, block := range f.Blocks {
			for i, instr := range block.Instructions {
				ansi.Dim.Printf("%%%-1d = ", i)
				fmt.Printf("%-30s ", instr.String())
				ansi.Dim.Printf(" %-30s", instr.Type().Name())
				fmt.Println()
			}
		}

		fmt.Println()
	})
}
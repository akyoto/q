package compiler

import (
	"fmt"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

// showSSA shows the SSA IR.
func showSSA(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		ansi.Yellow.Printf("%s:\n", f.UniqueName)

		for i, block := range f.Blocks {
			if i != 0 {
				fmt.Println("---")
			}

			for i, instr := range block.Instructions {
				ansi.Dim.Printf("%-4d", i)
				fmt.Printf("%-40s", instr.String())
				ansi.Cyan.Printf("%-30s", instr.Type().Name())
				ansi.Dim.Printf("%s\n", f.File.Bytes[instr.Start():instr.End()])
			}
		}

		fmt.Println()
	})
}
package compiler

import (
	"fmt"
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/go/color/ansi"
)

// showSSA shows the SSA IR.
func showSSA(functions iter.Seq[*core.Function]) {
	for f := range functions {
		ansi.Bold.Printf("%s:\n", f.UniqueName)

		for i, block := range f.Blocks {
			if i != 0 {
				fmt.Println("---")
			}

			for i, instr := range block.Instructions {
				fmt.Printf("t%d = %s\n", i, instr.String())
			}
		}
	}
}
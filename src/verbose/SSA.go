package verbose

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
	"unicode"

	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/ssa"
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

		tmpColor := ansi.Green
		tmpPrefix := "α"
		pointerToName := make(map[string]string, len(f.Steps))
		pointerToName[fmt.Sprintf("%p", ssa.Undefined)] = "?"

		for _, step := range f.Steps {
			pointerToName[fmt.Sprintf("%p", step.Value)] = fmt.Sprintf("%s%d", tmpPrefix, step.Index)
		}

		ansi.Yellow.Println(f.FullName + ":")

		for _, step := range f.Steps {
			_, isLabel := step.Value.(*codegen.Label)

			if isLabel {
				fmt.Printf("\n%s:", step.Value.String())

				for _, pre := range step.Block.Predecessors {
					ansi.Dim.Print(" ⇠ ")
					ansi.Dim.Print(pre)
				}

				fmt.Println()
				continue
			}

			ansi.Dim.Print("  ")
			ansi.Dim.Printf("%s%-2d", tmpPrefix, step.Index)
			ansi.Dim.Print(" = ")
			value := step.Value.String()
			_, isInt := step.Value.(*ssa.Int)

			if isInt {
				ansi.Cyan.Print(value)
			} else {
				pos := strings.Index(value, "0x")

				for {
					if pos == -1 {
						fmt.Print(value)
						break
					}

					end := strings.IndexFunc(value[pos+2:], func(r rune) bool {
						return !unicode.Is(unicode.Hex_Digit, r)
					})

					if end == -1 {
						end = len(value) - pos - 2
					}

					name := pointerToName[value[pos:pos+2+end]]
					fmt.Print(value[:pos])
					tmpColor.Printf("%s", name)
					value = value[pos+2+end:]
					pos = strings.Index(value, "0x")
				}
			}

			ansi.Dim.Printf(" %s %s ", step.Value.Type().Name(), step.Register)

			if step.Block.IsIdentified(step.Value) {
				ansi.Dim.Print("id: ")

				for identifier := range step.Block.IdentifiersFor(step.Value) {
					ansi.Dim.Printf("%s ", identifier)
				}
			}

			if len(step.Live) > 0 {
				ansi.Dim.Printf("live: ")

				slices.SortStableFunc(step.Live, func(a *codegen.Step, b *codegen.Step) int {
					return a.Index - b.Index
				})

				for _, live := range step.Live {
					ansi.Dim.Printf("%s%d ", tmpPrefix, live.Index)
				}
			}

			fmt.Println()
		}

		fmt.Println()
	})
}
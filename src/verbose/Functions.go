package verbose

import (
	"fmt"

	"git.urbach.dev/cli/q/src/core"
)

// Functions shows the list of live functions.
func Functions(root *core.Function) {
	root.EachDependency(make(map[*core.Function]bool), func(f *core.Function) {
		fmt.Println(f.FullName)
	})
}
package verbose

import (
	"fmt"

	"git.urbach.dev/cli/q/src/core"
)

// Functions shows the list of live functions.
func Functions(env *core.Environment) {
	for f := range env.LiveFunctions() {
		fmt.Println(f.FullName)
	}
}
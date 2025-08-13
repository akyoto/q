package verbose

import "git.urbach.dev/cli/q/src/core"

// Show shows additional information about the build.
func Show(env *core.Environment) {
	if env.Build.ShowSSA {
		if env.Build.ShowHeaders {
			Header(HeaderSSA)
		}

		SSA(env.Init)
	}

	if env.Build.ShowASM {
		if env.Build.ShowHeaders {
			Header(HeaderASM)
		}

		ASM(env.Init)
	}
}
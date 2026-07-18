package core

import "git.urbach.dev/cli/q/src/types"

// ReceiveEnum receives an enum from the scanner.
func (env *Environment) ReceiveEnum(enum *types.Enum) {
	pkg := env.AddPackage(enum.Package(), false)
	pkg.Enums[enum.Name()] = enum
}
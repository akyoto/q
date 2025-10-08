package core

import "git.urbach.dev/cli/q/src/types"

// ReceiveStruct receives a struct from the scanner.
func (env *Environment) ReceiveStruct(structure *types.Struct) {
	pkg := env.AddPackage(structure.Package, false)
	pkg.Structs[structure.Name()] = structure
}
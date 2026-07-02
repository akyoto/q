package core

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// TypeFromToken returns the type with the given token or `nil` if it doesn't exist.
func (env *Environment) TypeFromToken(typ token.Token, file *fs.File) (types.Type, error) {
	tokens := unwrapTypeToken(typ, file)
	return env.TypeFromTokens(tokens, file)
}
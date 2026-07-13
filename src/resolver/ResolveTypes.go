package resolver

import "git.urbach.dev/cli/q/src/core"

// ResolveTypes resolves all the type tokens in structs, globals and function parameters.
func ResolveTypes(env *core.Environment) error {
	err := parseStructs(env, env.Structs())

	if err != nil {
		return err
	}

	err = parseGlobals(env, env.Globals())

	if err != nil {
		return err
	}

	return parseParameters(env, env.Functions())
}
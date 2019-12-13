package build

import (
	"sort"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/stringutils/similarity"
)

// UnknownFunctionError produces an unknown function error
// and tries to guess which function the user was trying to type.
func (state *State) UnknownFunctionError(functionName string) error {
	knownFunctions := make([]string, 0, len(state.environment.Functions)+len(BuiltinFunctions))

	for builtin := range BuiltinFunctions {
		knownFunctions = append(knownFunctions, builtin)
	}

	for function := range state.environment.Functions {
		knownFunctions = append(knownFunctions, function)
	}

	// Suggest a function name based on the similarity to known functions
	sort.Slice(knownFunctions, func(a, b int) bool {
		aSimilarity := similarity.JaroWinkler(functionName, knownFunctions[a])
		bSimilarity := similarity.JaroWinkler(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.JaroWinkler(functionName, knownFunctions[0]) < 0.9 {
		return errors.New(&errors.UnknownFunction{Name: functionName})
	}

	return errors.New(&errors.UnknownFunction{
		Name:        functionName,
		CorrectName: knownFunctions[0],
	})
}

// UnknownPackageError produces an unknown package error
// and tries to guess which package the user was trying to type.
func (state *State) UnknownPackageError(pkgName string) error {
	if len(state.function.File.imports) == 0 {
		return &errors.UnknownPackage{Name: pkgName}
	}

	knownPackages := make([]string, 0, len(state.function.File.imports))

	for imp := range state.function.File.imports {
		knownPackages = append(knownPackages, imp)
	}

	// Suggest a package name based on the similarity to known functions
	sort.Slice(knownPackages, func(a, b int) bool {
		aSimilarity := similarity.JaroWinkler(pkgName, knownPackages[a])
		bSimilarity := similarity.JaroWinkler(pkgName, knownPackages[b])
		return aSimilarity > bSimilarity
	})

	if similarity.JaroWinkler(pkgName, knownPackages[0]) < 0.9 {
		return errors.New(&errors.UnknownPackage{Name: pkgName})
	}

	return errors.New(&errors.UnknownPackage{
		Name:        pkgName,
		CorrectName: knownPackages[0],
	})
}

// UnknownVariableError produces an unknown variable error
// and tries to guess which variable the user was trying to type.
func (state *State) UnknownVariableError(variableName string) error {
	knownVariables := []string{}

	state.scopes.Each(func(variable *Variable) {
		knownVariables = append(knownVariables, variable.Name)
	})

	if len(knownVariables) == 0 {
		return errors.New(&errors.UnknownVariable{Name: variableName})
	}

	// Suggest a variable name based on the similarity to known variables
	sort.Slice(knownVariables, func(a, b int) bool {
		aSimilarity := similarity.JaroWinkler(variableName, knownVariables[a])
		bSimilarity := similarity.JaroWinkler(variableName, knownVariables[b])
		return aSimilarity > bSimilarity
	})

	if similarity.JaroWinkler(variableName, knownVariables[0]) < 0.9 {
		return errors.New(&errors.UnknownVariable{Name: variableName})
	}

	return errors.New(&errors.UnknownVariable{
		Name:        variableName,
		CorrectName: knownVariables[0],
	})
}

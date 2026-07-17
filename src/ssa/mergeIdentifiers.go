package ssa

import (
	"maps"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

// mergeIdentifiers merges identifier mappings into the successor,
// inserting phi functions where values differ across predecessors.
func mergeIdentifiers(predecessor *Block, successor *Block) {
	if predecessor.Identifiers.After == nil {
		return
	}

	if successor.Identifiers.After == nil {
		successor.Identifiers.Before = make(map[string]Value, len(predecessor.Identifiers.After))
		successor.Identifiers.After = make(map[string]Value, len(predecessor.Identifiers.After))

		if len(successor.Predecessors) == 1 {
			maps.Copy(successor.Identifiers.Before, predecessor.Identifiers.After)
			maps.Copy(successor.Identifiers.After, predecessor.Identifiers.After)
			return
		}
	}

	var (
		modifiedStructs []string
		names           = collectIdentifierNames(predecessor, successor)
	)

	for _, name := range names {
		mergeIdentifier(predecessor, successor, name, &modifiedStructs)
	}

	// Structs that were modified in branches need to be recreated
	// to use the new Phi values as their arguments.
	for _, name := range modifiedStructs {
		structure := successor.Identifiers.Before[name].(*Struct)
		structType := types.Unwrap(structure.Typ).(*types.Struct)
		newStruct := &Struct{Typ: structure.Typ, Arguments: make(Arguments, len(structure.Arguments))}

		for i, field := range structType.Fields {
			newStruct.Arguments[i] = successor.Identifiers.Before[name+"."+field.Name]
		}

		successor.ReplaceIdentifier(name, structure, newStruct)
	}
}

// collectIdentifierNames returns all identifier names from both maps in deterministic order.
func collectIdentifierNames(predecessor *Block, successor *Block) []string {
	keys := make([]string, 0, max(len(predecessor.Identifiers.After), len(successor.Identifiers.After)))

	for name := range successor.Identifiers.Before {
		if !slices.Contains(keys, name) {
			keys = append(keys, name)
		}
	}

	for name := range predecessor.Identifiers.After {
		if !slices.Contains(keys, name) {
			keys = append(keys, name)
		}
	}

	slices.SortFunc(keys, func(a string, b string) int {
		return strings.Compare(b, a)
	})

	return keys
}

// mergeIdentifier merges a single identifier into the successor.
func mergeIdentifier(predecessor *Block, successor *Block, name string, modifiedStructs *[]string) {
	oldValue, oldExists := successor.Identifiers.Before[name]
	newValue, newExists := predecessor.Identifiers.After[name]

	switch {
	case oldExists:
		mergeOldIdentifier(successor, name, oldValue, newValue, newExists, modifiedStructs)
	case newExists:
		mergeNewIdentifier(successor, name, newValue)
	}
}

// mergeOldIdentifier handles the case where the successor already has a binding for this name.
func mergeOldIdentifier(successor *Block, name string, oldValue Value, newValue Value, newExists bool, modifiedStructs *[]string) {
	if oldValue == newValue {
		return
	}

	_, isStruct := oldValue.(*Struct)

	if isStruct {
		*modifiedStructs = append(*modifiedStructs, name)
		return
	}

	definedLocally := successor.Index(oldValue) != -1

	if definedLocally {
		phi, isPhi := oldValue.(*Phi)

		if isPhi {
			if newExists {
				phi.Arguments = append(phi.Arguments, newValue)
			} else {
				phi.Arguments = append(phi.Arguments, Undefined)
			}
		}

		return
	}

	phi := &Phi{
		Arguments: make([]Value, len(successor.Predecessors)-1, len(successor.Predecessors)),
		Typ:       oldValue.Type(),
	}

	for i := range phi.Arguments {
		phi.Arguments[i] = oldValue
	}

	successor.InsertAt(0, phi)
	successor.ReplaceIdentifier(name, oldValue, phi)

	if newExists {
		phi.Arguments = append(phi.Arguments, newValue)
	} else {
		phi.Arguments = append(phi.Arguments, Undefined)
	}
}

// mergeNewIdentifier handles the case where only the predecessor has a binding for this name.
func mergeNewIdentifier(successor *Block, name string, newValue Value) {
	phi := &Phi{
		Arguments: make([]Value, len(successor.Predecessors)-1, len(successor.Predecessors)),
		Typ:       newValue.Type(),
	}

	for i := range phi.Arguments {
		phi.Arguments[i] = Undefined
	}

	successor.InsertAt(0, phi)
	successor.ReplaceIdentifier(name, nil, phi)
	phi.Arguments = append(phi.Arguments, newValue)
}
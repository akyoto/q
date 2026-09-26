package ssa

import (
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/types"
)

// mergeIdentifiers merges identifier mappings into the successor,
// inserting phi functions where values differ across predecessors.
func mergeIdentifiers(predecessor *Block, successor *Block) {
	if predecessor.Identifiers.After.Raw() == nil {
		return
	}

	if successor.Identifiers.After.Raw() == nil {
		if len(successor.Predecessors) == 1 {
			after := predecessor.Identifiers.After.Raw()
			successor.Identifiers.Before.Share(after)
			successor.Identifiers.After.Share(after)
			return
		}

		fresh := make(map[string]Value, predecessor.Identifiers.After.Count())
		successor.Identifiers.Before.Share(fresh)
		successor.Identifiers.After.Share(fresh)
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
		value, _ := successor.Identifiers.Before.Get(name)
		structure := value.(*Struct)
		structType := types.Unwrap(structure.Typ).(*types.Struct)
		newStruct := &Struct{Typ: structure.Typ, Arguments: make(Arguments, len(structure.Arguments))}

		for i, field := range structType.Fields {
			fieldValue, _ := successor.Identifiers.Before.Get(name + "." + field.Name)
			newStruct.Arguments[i] = fieldValue
		}

		successor.ReplaceIdentifier(name, structure, newStruct)
	}
}

// collectIdentifierNames returns all identifier names from both maps in deterministic order.
func collectIdentifierNames(predecessor *Block, successor *Block) []string {
	keys := make([]string, 0, max(predecessor.Identifiers.After.Count(), successor.Identifiers.Before.Count()))

	for name := range successor.Identifiers.Before.Raw() {
		if !slices.Contains(keys, name) {
			keys = append(keys, name)
		}
	}

	for name := range predecessor.Identifiers.After.Raw() {
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
	oldValue, oldExists := successor.Identifiers.Before.Get(name)
	newValue, newExists := predecessor.Identifiers.After.Get(name)

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
		Name:      name,
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
		Name:      name,
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
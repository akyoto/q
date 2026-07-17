package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// deleteResources inserts delete calls for all resources.
func (f *Function) deleteResources(filter map[string]ssa.Value) {
	var (
		names       []string
		identifiers = f.Block().Identifiers.After
	)

	if len(identifiers) == 0 {
		return
	}

	for name := range identifiers {
		if filter != nil {
			_, inherited := filter[name]

			if inherited {
				continue
			}
		}

		if names == nil {
			names = make([]string, 0, len(identifiers))
		}

		names = append(names, name)
	}

	slices.SortFunc(names, func(a string, b string) int {
		aValue := identifiers[a]
		bValue := identifiers[b]

		aStruct, aIsStruct := aValue.(*ssa.Struct)

		if aIsStruct && len(aStruct.Arguments) > 0 {
			aValue = aStruct.Arguments[0]
		}

		bStruct, bIsStruct := bValue.(*ssa.Struct)

		if bIsStruct && len(bStruct.Arguments) > 0 {
			bValue = bStruct.Arguments[0]
		}

		for _, block := range f.IR.Blocks {
			aIndex := block.Index(aValue)
			bIndex := block.Index(bValue)

			switch {
			case aIndex != -1 && bIndex != -1:
				return bIndex - aIndex
			case aIndex != -1:
				return 1
			case bIndex != -1:
				return -1
			}
		}

		panic("non-deterministic memory deallocation order")
	})

	for _, name := range names {
		value := identifiers[name]

		if value == nil {
			continue
		}

		resource, isResource := value.Type().(*types.Resource)

		if !isResource {
			continue
		}

		_, isPointer := resource.Of.(*types.Pointer)
		_, isStruct := resource.Of.(*types.Struct)

		if !isPointer && !isStruct {
			continue
		}

		_, isParam := value.(*ssa.Parameter)

		if isParam {
			continue
		}

		f.delete(value)
	}
}
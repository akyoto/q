package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// deleteResources inserts delete calls for all resources.
func (f *Function) deleteResources() {
	identifiers := f.Block().IdentifiersAfter

	if len(identifiers) == 0 {
		return
	}

	names := make([]string, 0, len(identifiers))

	for k := range identifiers {
		names = append(names, k)
	}

	slices.SortFunc(names, func(a string, b string) int {
		aValue := identifiers[a]
		bValue := identifiers[b]

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

		return 0
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
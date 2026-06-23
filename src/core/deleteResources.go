package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// deleteResources inserts delete calls for all resources.
func (f *Function) deleteResources() {
	for _, value := range f.Block().IdentifiersAfter {
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
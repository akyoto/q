package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// checkResources checks for resources that were not deconstructed.
func (f *Function) checkResources() error {
	for exitBlock := range f.ExitBlocks {
		for _, value := range exitBlock.Identifiers {
			if value == nil {
				continue
			}

			_, isParam := value.(*ssa.Parameter)

			if isParam {
				continue
			}

			resource, isResource := value.Type().(*types.Resource)

			if !isResource {
				continue
			}

			return errors.New(&ResourceNotConsumed{TypeName: resource.Name()}, f.File, value.(ssa.HasSource).Start())
		}
	}

	return nil
}
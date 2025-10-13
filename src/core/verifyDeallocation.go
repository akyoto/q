package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// verifyDeallocation verifies that all resources have been deallocated.
func (f *Function) verifyDeallocation() error {
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

			phi, isPhi := value.(*ssa.Phi)

			if isPhi && phi.IsPartiallyUndefined() {
				return errors.NewAt(&ResourcePartiallyConsumed{TypeName: resource.Name()}, f.File, phi.FirstDefined().(errors.Source).Start())
			}

			return errors.NewAt(&ResourceNotConsumed{TypeName: resource.Name()}, f.File, value.(errors.Source).Start())
		}
	}

	return nil
}
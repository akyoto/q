package compiler

import (
	"iter"
	"slices"
	"strings"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
)

// findError returns compilation errors and creates a
// deterministic order in case of multiple errors.
func findError(functions iter.Seq[*core.Function]) error {
	var errs []error

	for f := range functions {
		if f.Err != nil {
			errs = append(errs, f.Err)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	if len(errs) == 1 {
		return errs[0]
	}

	slices.SortFunc(errs, func(errA error, errB error) int {
		a := errA.(*errors.FileError)
		b := errB.(*errors.FileError)
		diff := strings.Compare(a.File().Path, b.File().Path)

		if diff != 0 {
			return diff
		}

		lineA, columnA := a.LineColumn()
		lineB, columnB := b.LineColumn()

		if lineA == lineB {
			return columnA - columnB
		}

		return lineA - lineB
	})

	return &MultiError{Errors: errs}
}
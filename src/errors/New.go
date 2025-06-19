package errors

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// New generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func New(err error, file *fs.File, position token.Position) *FileError {
	return &FileError{
		err:      err,
		file:     file,
		position: position,
		stack:    stack(),
	}
}
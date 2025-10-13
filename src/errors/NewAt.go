package errors

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// NewAt is the same as New but with an exact position.
func NewAt(err error, file *fs.File, position token.Position) *FileError {
	return New(err, file, token.NewSource(position, position))
}
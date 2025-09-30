package core

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Global represents a variable that is visible in the entire package.
type Global struct {
	Name   string
	Typ    types.Type
	Tokens token.List
	File   *fs.File
}
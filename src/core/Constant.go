package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
)

// Constant registers a single value to be accessible under a descriptive name.
type Constant struct {
	File  *fs.File
	Value *expression.Expression
	Name  string
}
package core

import (
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// functionIdentity contains fields that identify the function.
type functionIdentity struct {
	File *fs.File
	Type *types.Function
	name string
	pkg  string
	body token.Source
}

// Body returns the function body.
func (f *functionIdentity) Body() token.List {
	return f.File.Tokens[f.body.Start():f.body.End()]
}

// IsExtern returns true if the function has no body.
func (f *functionIdentity) IsExtern() bool {
	return f.body.End() == 0
}

// Name returns the function name.
func (f *functionIdentity) Name() string {
	return f.name
}

// Package returns the package name.
func (f *functionIdentity) Package() string {
	return f.pkg
}

// SetBody sets the token range for the function body.
func (f *functionIdentity) SetBody(start int, end int) {
	f.body = token.NewSource(token.Position(start), token.Position(end))
}
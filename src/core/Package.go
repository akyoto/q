package core

import "git.urbach.dev/cli/q/src/types"

// Package represents a directory of functions.
type Package struct {
	Constants map[string]*Constant
	Functions map[string][]*Function
	Structs   map[string]*types.Struct
	Name      string
	IsExtern  bool
}
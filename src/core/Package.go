package core

import "git.urbach.dev/cli/q/src/types"

// Package represents a directory of functions.
type Package struct {
	Constants map[string]*Constant
	Enums     map[string]*types.Enum
	Functions map[string]*Function
	Structs   map[string]*types.Struct
	Globals   map[string]*Global
	Name      string
	IsExtern  bool
}
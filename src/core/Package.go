package core

// Package represents a directory of functions.
type Package struct {
	Constants map[string]*Constant
	Functions map[string]*Function
	Name      string
	IsExtern  bool
}
package core

// Package represents a directory of functions.
type Package struct {
	Name      string
	Functions map[string]*Function
	IsExtern  bool
}
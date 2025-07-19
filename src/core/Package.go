package core

// Package represents a directory of functions.
type Package struct {
	Functions map[string]*Function
	Name      string
	IsExtern  bool
}
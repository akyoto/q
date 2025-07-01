package asm

import "git.urbach.dev/cli/q/src/dll"

type compiler struct {
	code         []byte
	data         []byte
	dataLabels   map[string]Address
	labels       map[string]Address
	libraries    dll.List
	importsStart int
	deferred     []func()
}

func (c *compiler) Defer(call func()) {
	c.deferred = append(c.deferred, call)
}
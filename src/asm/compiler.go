package asm

import "git.urbach.dev/cli/q/src/dll"

type compiler struct {
	code                []byte
	data                []byte
	dataLabels          map[string]Address
	labels              map[string]Address
	libraries           dll.List
	importsStart        int
	deferred            map[int]func(int)
	deferredCodeChanges map[int]func(int) bool
}

func (c *compiler) Defer(address int, call func(int)) {
	c.deferred[address] = call
}

func (c *compiler) DeferCodeChange(address int, call func(int) bool) {
	c.deferredCodeChanges[address] = call
}
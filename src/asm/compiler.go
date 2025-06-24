package asm

type compiler struct {
	code       []byte
	data       []byte
	dataLabels map[string]Address
	labels     map[string]Address
	deferred   []func()
}

func (c *compiler) Defer(call func()) {
	c.deferred = append(c.deferred, call)
}
package asm

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
)

type compiler struct {
	patcher
	build        *config.Build
	data         []byte
	dataLabels   map[string]int
	libraries    dll.List
	importsStart int
}

// AddDataLabels adds the data labels to the existing labels.
// This can only run after the code size is known.
func (c *compiler) AddDataLabels() {
	x := exe.New(elf.HeaderEnd, c.build.FileAlign(), c.build.MemoryAlign(), c.build.Congruent(), c.code, c.data, nil)
	dataSectionOffset := x.Sections[1].MemoryOffset - x.Sections[0].MemoryOffset

	for dataLabel, address := range c.dataLabels {
		c.labels[dataLabel] = dataSectionOffset + address
	}

	if c.build.OS == config.Windows {
		c.importsStart = x.Sections[2].MemoryOffset - x.Sections[0].MemoryOffset
	}
}
package asm

import (
	"git.urbach.dev/cli/q/src/config"
	"git.urbach.dev/cli/q/src/dll"
	"git.urbach.dev/cli/q/src/elf"
	"git.urbach.dev/cli/q/src/exe"
	"git.urbach.dev/cli/q/src/macho"
	"git.urbach.dev/cli/q/src/pe"
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
	headerEnd := 0
	embedHeaders := false

	switch c.build.OS {
	case config.Linux:
		headerEnd = elf.HeaderEnd
	case config.Mac:
		headerEnd = macho.HeaderEnd
		embedHeaders = true
	case config.Windows:
		headerEnd = pe.HeaderEnd
	}

	x := exe.New(headerEnd, c.build.FileAlign(), c.build.MemoryAlign(), c.build.Congruent(), embedHeaders, c.code, c.data, nil)
	dataSectionOffset := x.Sections[1].MemoryOffset - x.Sections[0].MemoryOffset

	for dataLabel, address := range c.dataLabels {
		c.labels[dataLabel] = dataSectionOffset + address
	}

	if c.build.OS == config.Windows {
		c.importsStart = x.Sections[2].MemoryOffset - x.Sections[0].MemoryOffset
	}
}
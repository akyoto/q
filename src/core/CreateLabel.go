package core

import (
	"strconv"
	"strings"

	"git.urbach.dev/cli/q/src/asm"
)

// CreateLabel creates a label that is tied to this function by using a suffix.
func (f *Function) CreateLabel(prefix string, count counter) *asm.Label {
	tmp := strings.Builder{}
	tmp.WriteString(prefix)
	tmp.WriteString(" ")
	tmp.WriteString(strconv.FormatUint(uint64(count), 10))
	tmp.WriteString(" [")
	tmp.WriteString(f.UniqueName)
	tmp.WriteString("]")
	return &asm.Label{Name: tmp.String()}
}
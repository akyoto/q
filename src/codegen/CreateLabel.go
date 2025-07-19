package codegen

import (
	"strconv"
	"strings"
)

// CreateLabel creates a label that is tied to this function by using a suffix.
func (f *Function) CreateLabel(prefix string, count counter) string {
	tmp := strings.Builder{}
	tmp.WriteString(prefix)
	tmp.WriteString(" ")
	tmp.WriteString(strconv.FormatUint(uint64(count), 10))
	tmp.WriteString(" [")
	tmp.WriteString(f.FullName)
	tmp.WriteString("]")
	return tmp.String()
}
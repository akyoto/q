package elf

// SectionFlags defines flags for sections.
type SectionFlags int64

const (
	SectionFlagsWritable   SectionFlags = 1 << 0
	SectionFlagsAllocate   SectionFlags = 1 << 1
	SectionFlagsExecutable SectionFlags = 1 << 2
	SectionFlagsStrings    SectionFlags = 1 << 5
	SectionFlagsTLS        SectionFlags = 1 << 10
)
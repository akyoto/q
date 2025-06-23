package elf

// SectionType defines the type of the section.
type SectionType int32

const (
	SectionTypeNULL     SectionType = 0
	SectionTypePROGBITS SectionType = 1
	SectionTypeSYMTAB   SectionType = 2
	SectionTypeSTRTAB   SectionType = 3
	SectionTypeRELA     SectionType = 4
	SectionTypeHASH     SectionType = 5
	SectionTypeDYNAMIC  SectionType = 6
	SectionTypeNOTE     SectionType = 7
	SectionTypeNOBITS   SectionType = 8
	SectionTypeREL      SectionType = 9
	SectionTypeSHLIB    SectionType = 10
	SectionTypeDYNSYM   SectionType = 11
)
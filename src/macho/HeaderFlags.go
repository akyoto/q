package macho

type HeaderFlags uint32

const (
	FlagNoUndefs              HeaderFlags = 0x1
	FlagIncrLink              HeaderFlags = 0x2
	FlagDyldLink              HeaderFlags = 0x4
	FlagBindAtLoad            HeaderFlags = 0x8
	FlagPrebound              HeaderFlags = 0x10
	FlagSplitSegs             HeaderFlags = 0x20
	FlagLazyInit              HeaderFlags = 0x40
	FlagTwoLevel              HeaderFlags = 0x80
	FlagForceFlat             HeaderFlags = 0x100
	FlagNoMultiDefs           HeaderFlags = 0x200
	FlagNoFixPrebinding       HeaderFlags = 0x400
	FlagPrebindable           HeaderFlags = 0x800
	FlagAllModsBound          HeaderFlags = 0x1000
	FlagSubsectionsViaSymbols HeaderFlags = 0x2000
	FlagCanonical             HeaderFlags = 0x4000
	FlagWeakDefines           HeaderFlags = 0x8000
	FlagBindsToWeak           HeaderFlags = 0x10000
	FlagAllowStackExecution   HeaderFlags = 0x20000
	FlagRootSafe              HeaderFlags = 0x40000
	FlagSetuidSafe            HeaderFlags = 0x80000
	FlagNoReexportedDylibs    HeaderFlags = 0x100000
	FlagPIE                   HeaderFlags = 0x200000
	FlagDeadStrippableDylib   HeaderFlags = 0x400000
	FlagHasTLVDescriptors     HeaderFlags = 0x800000
	FlagNoHeapExecution       HeaderFlags = 0x1000000
	FlagAppExtensionSafe      HeaderFlags = 0x2000000
)
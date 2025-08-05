package macho

type LoadCommand uint32

const (
	LcSegment            LoadCommand = 0x1
	LcSymtab             LoadCommand = 0x2
	LcThread             LoadCommand = 0x4
	LcUnixthread         LoadCommand = 0x5
	LcDysymtab           LoadCommand = 0xB
	LcLoadDylib          LoadCommand = 0xC
	LcIdDylib            LoadCommand = 0xD
	LcLoadDylinker       LoadCommand = 0xE
	LcIdDylinker         LoadCommand = 0xF
	LcSegment64          LoadCommand = 0x19
	LcUuid               LoadCommand = 0x1B
	LcRpath              LoadCommand = 0x8000001C
	LcCodeSignature      LoadCommand = 0x1D
	LcSegmentSplitInfo   LoadCommand = 0x1E
	LcEncryptionInfo     LoadCommand = 0x21
	LcDyldInfo           LoadCommand = 0x22
	LcDyldInfoOnly       LoadCommand = 0x80000022
	LcVersionMinMacosx   LoadCommand = 0x24
	LcVersionMinIphoneos LoadCommand = 0x25
	LcFunctionStarts     LoadCommand = 0x26
	LcDyldEnvironment    LoadCommand = 0x27
	LcMain               LoadCommand = 0x80000028
	LcDataInCode         LoadCommand = 0x29
	LcSourceVersion      LoadCommand = 0x2A
	LcDylibCodeSignDrs   LoadCommand = 0x2B
	LcEncryptionInfo64   LoadCommand = 0x2C
	LcVersionMinTvos     LoadCommand = 0x2F
	LcVersionMinWatchos  LoadCommand = 0x30
	LcBuildVersion       LoadCommand = 0x32
	LcDyldChainedFixups  LoadCommand = 0x80000034
)
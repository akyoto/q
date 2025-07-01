package pe

const DLLImportSize = 20

type DLLImport struct {
	RvaFunctionNameList    uint32
	TimeDateStamp          uint32
	ForwarderChain         uint32
	RvaModuleName          uint32
	RvaFunctionAddressList uint32
}
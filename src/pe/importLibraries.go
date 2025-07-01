package pe

import "git.urbach.dev/cli/q/src/dll"

// importLibraries generates the import address table which contains the addresses of functions imported from DLLs.
func importLibraries(dlls dll.List, importsStart int) ([]uint64, []byte, []DLLImport, int) {
	imports := make([]uint64, 0)
	dllData := make([]byte, 0)
	dllImports := []DLLImport{}

	for _, library := range dlls {
		functionsStart := len(imports) * 8
		dllNamePos := len(dllData)
		dllData = append(dllData, library.Name...)
		dllData = append(dllData, ".dll"...)
		dllData = append(dllData, 0x00)

		dllImports = append(dllImports, DLLImport{
			RvaFunctionNameList:    uint32(importsStart + functionsStart),
			TimeDateStamp:          0,
			ForwarderChain:         0,
			RvaModuleName:          uint32(dllNamePos),
			RvaFunctionAddressList: uint32(importsStart + functionsStart),
		})

		for _, fn := range library.Functions {
			if len(dllData)&1 != 0 {
				dllData = append(dllData, 0x00) // align the next entry on an even boundary
			}

			offset := len(dllData)
			dllData = append(dllData, 0x00, 0x00)
			dllData = append(dllData, fn...)
			dllData = append(dllData, 0x00)

			imports = append(imports, uint64(offset))
		}

		imports = append(imports, 0)
	}

	dllDataStart := importsStart + len(imports)*8

	for i := range imports {
		if imports[i] == 0 {
			continue
		}

		imports[i] += uint64(dllDataStart)
	}

	for i := range dllImports {
		dllImports[i].RvaModuleName += uint32(dllDataStart)
	}

	// a zeroed structure marks the end of the list
	dllImports = append(dllImports, DLLImport{})

	return imports, dllData, dllImports, dllDataStart
}
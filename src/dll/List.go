package dll

import "slices"

// List is a slice of DLLs.
type List []Library

// Append adds a function for the given DLL if it doesn't exist yet.
func (list List) Append(dllName string, funcName string) List {
	for i, dll := range list {
		if dll.Name != dllName {
			continue
		}

		if slices.Contains(dll.Functions, funcName) {
			return list
		}

		list[i].Functions = append(list[i].Functions, funcName)
		return list
	}

	return append(list, Library{Name: dllName, Functions: []string{funcName}})
}

// Contains returns true if the library exists.
func (list List) Contains(dllName string) bool {
	for _, dll := range list {
		if dll.Name == dllName {
			return true
		}
	}

	return false
}

// Index returns the position of the given function name.
func (list List) Index(dllName string, funcName string) int {
	index := 0

	for _, dll := range list {
		if dll.Name != dllName {
			index += len(dll.Functions) + 1
			continue
		}

		for i, fn := range dll.Functions {
			if fn == funcName {
				return index + i
			}
		}
	}

	return -1
}
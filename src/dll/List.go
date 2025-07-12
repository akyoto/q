package dll

import (
	"iter"
	"slices"
)

// List is a slice of DLLs.
type List struct {
	libs []Library
}

// All returns an iterator over all libraries.
func (list *List) All() iter.Seq[Library] {
	return func(yield func(Library) bool) {
		for _, lib := range list.libs {
			if !yield(lib) {
				return
			}
		}
	}
}

// Append adds a function for the given DLL if it doesn't exist yet.
func (list *List) Append(dllName string, funcName string) {
	for i, dll := range list.libs {
		if dll.Name != dllName {
			continue
		}

		if slices.Contains(dll.Functions, funcName) {
			return
		}

		list.libs[i].Functions = append(list.libs[i].Functions, funcName)
		return
	}

	list.libs = append(list.libs, Library{Name: dllName, Functions: []string{funcName}})
}

// Contains returns true if the library exists.
func (list *List) Contains(dllName string) bool {
	for _, dll := range list.libs {
		if dll.Name == dllName {
			return true
		}
	}

	return false
}

// Index returns the position of the given function name.
func (list *List) Index(dllName string, funcName string) int {
	index := 0

	for _, dll := range list.libs {
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
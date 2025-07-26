package core

// EachDependency recursively finds all the calls to other functions.
// It avoids calling the same function twice with the help of a hashmap.
func (f *Function) EachDependency(traversed map[*Function]bool, call func(*Function)) {
	call(f)
	traversed[f] = true

	for dep := range f.Dependencies.All() {
		if traversed[dep] {
			continue
		}

		dep.EachDependency(traversed, call)
	}
}
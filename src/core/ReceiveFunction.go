package core

// ReceiveFunction receives a function from the scanner.
func (env *Environment) ReceiveFunction(f *Function) {
	f.Env = env
	f.CPU = env.Build.CPU()
	pkg := env.AddPackage(f.Package(), f.IsExtern())
	env.NumFunctions++

	existing := pkg.Functions[f.Name()]

	if existing == nil {
		pkg.Functions[f.Name()] = f
		return
	}

	for existing.Next != nil {
		existing = existing.Next
	}

	existing.Next = f
	f.Previous = existing
}
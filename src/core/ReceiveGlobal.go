package core

// ReceiveGlobal receives a global from the scanner.
func (env *Environment) ReceiveGlobal(global *Global) {
	pkg := env.AddPackage(global.File.Package, false)
	pkg.Globals[global.Name] = global
}
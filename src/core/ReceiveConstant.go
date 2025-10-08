package core

// ReceiveConstant receives a constant from the scanner.
func (env *Environment) ReceiveConstant(constant *Constant) {
	pkg := env.AddPackage(constant.File.Package, false)
	pkg.Constants[constant.Name] = constant
}
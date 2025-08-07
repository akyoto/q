package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileASTNode compiles a node in the AST.
func (f *Function) compileASTNode(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Assert:
		return f.compileAssert(node)

	case *ast.Assign:
		return f.compileAssign(node)

	case *ast.Call:
		return f.compileCall(node)

	case *ast.Define:
		return f.compileDefinition(node)

	case *ast.If:
		return f.compileIf(node)

	case *ast.Loop:
		return f.compileLoop(node)

	case *ast.Return:
		return f.compileReturn(node)

	case *ast.Switch:
		return f.compileSwitch(node)

	default:
		panic("unknown AST type")
	}
}
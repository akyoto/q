package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileLoop compiles an endless loop.
func (f *Function) compileLoop(loop *ast.Loop) error {
	return f.loop(nil, loop.Body)
}
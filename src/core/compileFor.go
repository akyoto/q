package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileFor compiles a loop with an exit condition.
func (f *Function) compileFor(loop *ast.For) error {
	return f.loop(loop.Head, loop.Body)
}
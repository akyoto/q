package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileGo compiles a call that runs asynchronously in a separate thread.
func (f *Function) compileGo(g *ast.Go) error {
	threadFunc, err := f.evaluate(g.Call.Children[0])

	if err != nil {
		return err
	}

	f.Append(threadFunc)
	createThread := f.Env.Function("thread", "create")

	f.Append(&ssa.Call{
		Func:      &ssa.Function{Typ: createThread.Type, FunctionRef: createThread},
		Arguments: []ssa.Value{threadFunc},
	})

	f.Dependencies.Add(createThread)
	return nil
}
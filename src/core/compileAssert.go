package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileAssert compiles an assertion.
func (f *Function) compileAssert(assert *ast.Assert) error {
	f.Count.Assert++
	thenLabel := f.CreateLabel("assert.then", f.Count.Assert)
	elseLabel := f.CreateLabel("assert.else", f.Count.Assert)
	thenBlock := ssa.NewBlock(thenLabel)
	elseBlock := ssa.NewBlock(elseLabel)
	f.Block().AddSuccessor(thenBlock)
	f.Block().AddSuccessor(elseBlock)
	err := f.compileCondition(assert.Condition, thenBlock, elseBlock)

	if err != nil {
		return err
	}

	f.AddBlock(elseBlock)
	crash := f.Env.Function("run", "crash")

	elseBlock.Append(&ssa.Call{Func: &ssa.Function{
		Typ:         crash.Type,
		FunctionRef: crash,
	}})

	f.Dependencies.Add(crash)
	f.AddBlock(thenBlock)
	return nil
}
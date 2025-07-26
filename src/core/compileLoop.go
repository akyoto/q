package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileLoop compiles an endless loop.
func (f *Function) compileLoop(loop *ast.Loop) error {
	f.Count.Loop++
	bodyLabel := f.CreateLabel("loop.body", f.Count.Loop)
	exitLabel := f.CreateLabel("loop.exit", f.Count.Loop)
	beforeLoop := f.Block()
	loopBody := ssa.NewBlock(bodyLabel)

	beforeLoop.AddSuccessor(loopBody)
	loopBlockIndex := len(f.Blocks)
	f.AddBlock(loopBody)

	// Compile loop body
	err := f.compileAST(loop.Body)

	if err != nil {
		return err
	}

	loopBlocks := f.Blocks[loopBlockIndex:len(f.Blocks)]

	// Find identifiers defined outside the loop that are modified within the loop
	modified := set.Ordered[string]{}

	for _, block := range loopBlocks {
		for name := range block.Identifiers {
			_, existedBeforeLoop := beforeLoop.FindIdentifier(name)

			if existedBeforeLoop {
				modified.Add(name)
			}
		}
	}

	// Insert phi functions
	for identifier := range modified.All() {
		oldValue, _ := beforeLoop.FindIdentifier(identifier)
		newValue, _ := f.Block().FindIdentifier(identifier)
		phi := &ssa.Phi{Arguments: []ssa.Value{oldValue, newValue}, Typ: oldValue.Type()}

		for _, block := range loopBlocks {
			for _, instr := range block.Instructions {
				instr.Replace(oldValue, phi)
			}
		}

		loopBody.InsertAt(phi, 0)
	}

	f.Append(&ssa.Jump{To: loopBody})
	f.Block().AddSuccessor(loopBody)
	loopExit := ssa.NewBlock(exitLabel)
	f.AddBlock(loopExit)
	return nil
}
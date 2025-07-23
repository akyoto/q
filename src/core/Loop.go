package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// Loop compiles an endless loop.
func (f *Function) Loop(tokens token.List) error {
	blockStart, blockEnd, err := f.block(tokens)

	if err != nil {
		return err
	}

	f.Count.Loop++
	bodyLabel := f.CreateLabel("loop body", f.Count.Loop)
	exitLabel := f.CreateLabel("loop exit", f.Count.Loop)
	beforeLoop := f.Block()
	loopBody := ssa.NewBlock(bodyLabel)
	loopExit := ssa.NewBlock(exitLabel)
	beforeLoop.AddSuccessor(loopBody)
	beforeLoop.AddSuccessor(loopExit)
	loopBlockIndex := len(f.Blocks)
	f.AddBlock(loopBody)

	// Compile loop body
	err = f.compileTokens(tokens[blockStart+1 : blockEnd])

	if err != nil {
		return err
	}

	// Insert phi functions
	loopBlocks := f.Blocks[loopBlockIndex:len(f.Blocks)]

	for _, block := range loopBlocks {
		for name, newValue := range block.Identifiers {
			oldValue := beforeLoop.Identifiers[name]

			if oldValue != newValue {
				phi := &ssa.Phi{Arguments: []ssa.Value{oldValue, newValue}}

				for _, user := range oldValue.Users() {
					phi.AddUser(user)
				}

				for _, block := range loopBlocks {
					for _, instr := range block.Instructions {
						instr.Replace(oldValue, phi)
						liveness, hasLiveness := instr.(ssa.HasLiveness)

						if hasLiveness {
							liveness.ReplaceUser(oldValue, phi)
						}
					}
				}

				loopBody.InsertAt(phi, 0)
			}
		}
	}

	f.Append(&ssa.Jump{To: loopBody})
	f.AddBlock(loopExit)
	return nil
}
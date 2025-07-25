package core

import (
	"git.urbach.dev/cli/q/src/set"
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
	loopBody.AddSuccessor(loopExit)
	loopBlockIndex := len(f.Blocks)
	f.AddBlock(loopBody)

	// Compile loop body
	err = f.compileTokens(tokens[blockStart+1 : blockEnd])

	if err != nil {
		return err
	}

	loopBlocks := f.Blocks[loopBlockIndex:len(f.Blocks)]

	// Find identifiers defined outside the loop that are modified within the loop
	modified := set.Ordered[string]{}

	for _, block := range loopBlocks {
		for name := range block.Identifiers {
			_, existedBeforeLoop := beforeLoop.Identifiers[name]

			if existedBeforeLoop {
				modified.Add(name)
			}
		}
	}

	// Insert phi functions
	for identifier := range modified.All() {
		traversed := make(map[*ssa.Block]bool)
		oldValue, appended := beforeLoop.LookupIdentifier(identifier, traversed)

		if !appended {
			beforeLoop.Identify(identifier, oldValue)
			beforeLoop.Append(oldValue)
		}

		clear(traversed)
		newValue, appended := f.Block().LookupIdentifier(identifier, traversed)

		if !appended {
			f.Block().Identify(identifier, newValue)
			f.Block().Append(newValue)
		}

		phi := &ssa.Phi{Arguments: []ssa.Value{newValue, oldValue}}

		for _, block := range loopBlocks {
			for _, instr := range block.Instructions {
				instr.Replace(oldValue, phi)
			}
		}

		loopBody.InsertAt(phi, 0)
	}

	f.Append(&ssa.Jump{To: loopBody})
	f.AddBlock(loopExit)
	return nil
}
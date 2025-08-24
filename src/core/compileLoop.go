package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// compileLoop compiles an endless loop.
func (f *Function) compileLoop(loop *ast.Loop) error {
	f.Count.Loop++
	headLabel := f.CreateLabel("loop.head", f.Count.Loop)
	exitLabel := f.CreateLabel("loop.exit", f.Count.Loop)
	beforeLoop := f.Block()
	loopHead := ssa.NewBlock(headLabel)
	loopExit := ssa.NewBlock(exitLabel)
	loopBlockIndex := len(f.Blocks)

	if loop.Head != nil {
		// Before the loop starts, we evaluate the lower limit
		// and identify it as the loop counter.
		name, from, to := f.parseLoopHeader(loop.Head)

		if from == nil {
			return errors.New(InvalidLoopHeader, f.File, loop.Head.Source().StartPos)
		}

		fromValue, err := f.evaluate(from)

		if err != nil {
			return err
		}

		if f.Block().IsIdentified(fromValue) {
			fromValue = f.copy(fromValue, from.Source())
		}

		beforeLoop.Identify(name, fromValue)
		beforeLoop.Append(&ssa.Jump{To: loopHead})
		beforeLoop.AddSuccessor(loopHead)

		// Loop starts, this is the jump target for new iterations.
		// The upper limit is recalculated on every iteration.
		// We check that the condition to jump to the loop body is true,
		// otherwise we jump to the loop exit.
		f.AddBlock(loopHead)
		toValue, err := f.evaluate(to)

		if err != nil {
			return err
		}

		condition := f.Append(&ssa.BinaryOp{
			Op:    token.Less,
			Left:  fromValue,
			Right: toValue,
		})

		bodyLabel := f.CreateLabel("loop.body", f.Count.Loop)
		bodyBlock := ssa.NewBlock(bodyLabel)

		f.Append(&ssa.Branch{
			Condition: condition,
			Then:      bodyBlock,
			Else:      loopExit,
		})

		loopHead.AddSuccessor(bodyBlock)
		loopHead.AddSuccessor(loopExit)

		// Loop condition is true from now on so we'll
		// execute the code inside the loop body.
		f.AddBlock(bodyBlock)
		err = f.compileAST(loop.Body)

		if err != nil {
			return err
		}

		one := f.Append(&ssa.Int{Int: 1})

		nextIteration := f.Append(&ssa.BinaryOp{
			Op:    token.Add,
			Left:  fromValue,
			Right: one,
		})

		f.Block().Identify(name, nextIteration)
	} else {
		beforeLoop.Append(&ssa.Jump{To: loopHead})
		beforeLoop.AddSuccessor(loopHead)
		f.AddBlock(loopHead)

		// For infinite loops, there are no conditions to check,
		// we can simply process the loop body.
		err := f.compileAST(loop.Body)

		if err != nil {
			return err
		}
	}

	// The initial compilation of the loop body does not know
	// that the code is repeated in a loop. Therefore, we need
	// to find identifiers that were both defined outside the loop
	// and modified within the loop. For these identifiers,
	// we create phi functions at the top of the loop head.
	loopBlocks := f.Blocks[loopBlockIndex:len(f.Blocks)]
	modified := set.Ordered[string]{}

	for _, block := range loopBlocks {
		if block.Loop != nil {
			continue
		}

		block.Loop = loopHead

		for name := range block.Identifiers {
			_, existedBeforeLoop := beforeLoop.FindIdentifier(name)

			if existedBeforeLoop {
				modified.Add(name)
			}
		}
	}

	// Insert phi functions that capture both the value
	// outside of the loop and the modification within it.
	// We initially only knew about the value outside of the loop,
	// so we need to replace all of its occurrences in the loop blocks
	// with the new phi function.
	replacements := make(map[ssa.Value]*ssa.Phi, modified.Count())

	for identifier := range modified.All() {
		oldValue, _ := beforeLoop.FindIdentifier(identifier)
		newValue, _ := f.Block().FindIdentifier(identifier)
		phi := &ssa.Phi{Arguments: []ssa.Value{oldValue, newValue}, Typ: oldValue.Type()}
		replacement, exists := replacements[oldValue]

		if exists && replacement.Equals(phi) {
			loopHead.Identify(identifier, replacement)
			continue
		}

		replacements[oldValue] = phi

		for _, block := range loopBlocks {
			for _, instr := range block.Instructions {
				instr.Replace(oldValue, phi)
			}
		}

		loopHead.InsertAt(phi, 0)
		loopHead.Identify(identifier, phi)
	}

	// Jump back to the loop head.
	f.Append(&ssa.Jump{To: loopHead})
	f.Block().AddSuccessor(loopHead)
	f.AddBlock(loopExit)
	return nil
}
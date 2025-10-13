package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// compileLoop compiles an endless loop.
func (f *Function) compileLoop(node *ast.Loop) error {
	f.Count.Loop++
	headLabel := f.CreateLabel("loop.head", f.Count.Loop)
	exitLabel := f.CreateLabel("loop.exit", f.Count.Loop)
	beforeLoop := f.Block()
	loopHead := ssa.NewBlock(headLabel)
	loopExit := ssa.NewBlock(exitLabel)
	loopBlockIndex := len(f.Blocks)

	loop := &Loop{
		Head: loopHead,
		Exit: loopExit,
	}

	if node.Head != nil {
		// Before the loop starts, we evaluate the lower limit
		// and identify it as the loop counter.
		name, from, to := f.parseLoopHeader(node.Head)

		if from == nil {
			return errors.New(InvalidLoopHeader, f.File, node.Head.Source())
		}

		loop.IteratorName = name
		fromValue, err := f.evaluateRight(from)

		if err != nil {
			return err
		}

		if f.Block().IsIdentified(fromValue) {
			fromValue = f.copy(fromValue, from.Source())
		}

		loop.FromValue = fromValue
		beforeLoop.Identify(name, fromValue)
		f.jump(loopHead)

		// Loop starts, this is the jump target for new iterations.
		// The upper limit is recalculated on every iteration.
		// We check that the condition to jump to the loop body is true,
		// otherwise we jump to the loop exit.
		f.AddBlock(loopHead)
		toValue, err := f.evaluateRight(to)

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

		// Loop condition is true from now on so we'll
		// execute the code inside the loop body.
		f.AddBlock(bodyBlock)
		f.loopStack.Push(loop)
		err = f.compileAST(node.Body)

		if err != nil {
			return err
		}

		f.loopStack.Pop()
	} else {
		f.jump(loopHead)
		f.AddBlock(loopHead)

		// For infinite loops, there are no conditions to check,
		// we can simply process the loop body.
		f.loopStack.Push(loop)
		err := f.compileAST(node.Body)

		if err != nil {
			return err
		}

		f.loopStack.Pop()
	}

	// Jump back to the loop head.
	f.loopNext(loop)

	// The initial compilation of the loop body does not know
	// that the code is repeated in a loop. Therefore, we need
	// to find identifiers that were both defined outside the loop
	// and modified within the loop. For these identifiers,
	// we created Phi functions at the top of the loop head.
	// All that's left to do is to replace all the occurrences
	// of the old values with their new Phi in the loop blocks.
	loopBlocks := f.Blocks[loopBlockIndex:len(f.Blocks)]

	for _, block := range loopBlocks {
		if block.Loop != nil {
			continue
		}

		block.Loop = loopHead

		for phi := range loopHead.Phis {
			oldValue := phi.Arguments[0]

			if oldValue == ssa.Undefined {
				continue
			}

			for _, instr := range block.Instructions {
				if instr == phi {
					continue
				}

				instr.Replace(oldValue, phi)
			}
		}
	}

	if node.Head != nil {
		loopHead.AddSuccessor(loopExit)
	}

	f.AddBlock(loopExit)
	return nil
}
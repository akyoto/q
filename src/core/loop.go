package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// loop provides a generic way to create loops.
func (f *Function) loop(head *expression.Expression, body ast.AST) error {
	f.Count.Loop++
	bodyLabel := f.CreateLabel("loop.body", f.Count.Loop)
	exitLabel := f.CreateLabel("loop.exit", f.Count.Loop)
	beforeLoop := f.Block()
	loopBody := ssa.NewBlock(bodyLabel)
	loopExit := ssa.NewBlock(exitLabel)
	loopBlockIndex := len(f.Blocks)

	if head != nil {
		name, from, to := f.parseLoopHeader(head)

		if from == nil {
			return errors.New(InvalidLoopHeader, f.File, head.Source().StartPos)
		}

		fromValue, err := f.evaluate(from)

		if err != nil {
			return err
		}

		beforeLoop.Identify(name, fromValue)
		beforeLoop.Append(&ssa.Jump{To: loopBody})
		beforeLoop.AddSuccessor(loopBody)
		f.AddBlock(loopBody)

		thenBlock := ssa.NewBlock(f.CreateLabel("loop.then", f.Count.Loop))
		toValue, err := f.evaluate(to)

		if err != nil {
			return err
		}

		condition := f.Append(&ssa.BinaryOp{
			Left:  fromValue,
			Right: toValue,
			Op:    token.Less,
		})

		f.Append(&ssa.Branch{
			Condition: condition,
			Then:      thenBlock,
			Else:      loopExit,
		})

		loopBody.AddSuccessor(thenBlock)
		loopBody.AddSuccessor(loopExit)
		f.AddBlock(thenBlock)
		err = f.compileAST(body)

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
		beforeLoop.Append(&ssa.Jump{To: loopBody})
		beforeLoop.AddSuccessor(loopBody)
		f.AddBlock(loopBody)
		err := f.compileAST(body)

		if err != nil {
			return err
		}
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

	replacements := make(map[ssa.Value]*ssa.Phi, modified.Count())

	// Insert phi functions
	for identifier := range modified.All() {
		oldValue, _ := beforeLoop.FindIdentifier(identifier)
		newValue, _ := f.Block().FindIdentifier(identifier)
		phi := &ssa.Phi{Arguments: []ssa.Value{oldValue, newValue}, Typ: oldValue.Type()}
		replacement, exists := replacements[oldValue]

		if exists && replacement.Equals(phi) {
			loopBody.Identify(identifier, replacement)
			continue
		}

		replacements[oldValue] = phi

		for _, block := range loopBlocks {
			for _, instr := range block.Instructions {
				instr.Replace(oldValue, phi)
			}
		}

		loopBody.InsertAt(phi, 0)
		loopBody.Identify(identifier, phi)
	}

	f.Append(&ssa.Jump{To: loopBody})
	f.Block().AddSuccessor(loopBody)
	f.AddBlock(loopExit)
	return nil
}
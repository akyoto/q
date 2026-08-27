package codegen_test

import (
	"maps"
	"testing"

	"git.urbach.dev/cli/q/src/asm"
	"git.urbach.dev/cli/q/src/codegen"
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/go/assert"
)

const (
	r0 = cpu.Register(0)
	r1 = cpu.Register(1)
	r2 = cpu.Register(2)
	r3 = cpu.Register(3)
	r4 = cpu.Register(4)
	r5 = cpu.Register(5)
)

func TestScheduleMovesEmpty(t *testing.T) {
	scheduled, ok := codegen.ScheduleMoves(nil, []cpu.Register{r3})
	assert.True(t, ok)
	assert.Equal(t, 0, len(scheduled))
}

func TestScheduleMovesSingle(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
	}

	expected := []*asm.Move{
		mv(r1, r0),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
}

func TestScheduleMovesIndependent(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r3, r2),
	}

	expected := []*asm.Move{
		mv(r1, r0),
		mv(r3, r2),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r4})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
}

func TestScheduleMovesChain(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r2, r1),
	}

	expected := []*asm.Move{
		mv(r2, r1),
		mv(r1, r0),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
	verify(t, moves, []cpu.Register{r3})
}

func TestScheduleMovesNoOp(t *testing.T) {
	moves := []*asm.Move{
		mv(r0, r0),
		mv(r2, r1),
	}

	expected := []*asm.Move{
		mv(r2, r1),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
}

func TestScheduleMovesTwoCycle(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r0, r1),
	}

	expected := []*asm.Move{
		mv(r3, r0),
		mv(r0, r1),
		mv(r1, r3),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
	verify(t, moves, []cpu.Register{r3})
}

func TestScheduleMovesThreeCycle(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r2, r1),
		mv(r0, r2),
	}

	expected := []*asm.Move{
		mv(r3, r0),
		mv(r0, r2),
		mv(r2, r1),
		mv(r1, r3),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
	verify(t, moves, []cpu.Register{r3})
}

func TestScheduleMovesFourCycle(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r2, r1),
		mv(r3, r2),
		mv(r0, r3),
	}

	expected := []*asm.Move{
		mv(r4, r0),
		mv(r0, r3),
		mv(r3, r2),
		mv(r2, r1),
		mv(r1, r4),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r4})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
	verify(t, moves, []cpu.Register{r4})
}

func TestScheduleMovesNoFreeRegister(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r0, r1),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, nil)
	assert.False(t, ok)
	assert.Nil(t, scheduled)
}

func TestScheduleMovesNoFreeRegisterNoCycle(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
	}

	expected := []*asm.Move{
		mv(r1, r0),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, nil)
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
}

func TestScheduleMovesSkipsUsedFree(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r0, r1),
	}

	expected := []*asm.Move{
		mv(r3, r0),
		mv(r0, r1),
		mv(r1, r3),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r0, r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
	verify(t, moves, []cpu.Register{r0, r3})
}

func TestScheduleMovesMultipleCycles(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r2, r1),
		mv(r0, r2),
		mv(r4, r3),
		mv(r3, r4),
	}

	free := []cpu.Register{r5}
	verify(t, moves, free)
}

func TestScheduleMovesMixed(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r3, r2),
		mv(r4, r3),
		mv(r2, r4),
	}

	free := []cpu.Register{r5}
	verify(t, moves, free)
}

func TestScheduleMovesDestinationNotRead(t *testing.T) {
	moves := []*asm.Move{
		mv(r1, r0),
		mv(r2, r1),
	}

	expected := []*asm.Move{
		mv(r2, r1),
		mv(r1, r0),
	}

	scheduled, ok := codegen.ScheduleMoves(moves, []cpu.Register{r3})
	assert.True(t, ok)
	assert.DeepEqual(t, scheduled, expected)
}

// mv creates a move from source to destination.
func mv(destination cpu.Register, source cpu.Register) *asm.Move {
	return &asm.Move{Destination: destination, Source: source}
}

// verify checks that the scheduled moves, when executed in order,
// produce the same final register values as executing the original
// moves simultaneously.
func verify(t *testing.T, moves []*asm.Move, free []cpu.Register) {
	scheduled, ok := codegen.ScheduleMoves(moves, free)
	assert.True(t, ok)
	registers := map[cpu.Register]struct{}{}

	for _, move := range moves {
		registers[move.Source] = struct{}{}
		registers[move.Destination] = struct{}{}
	}

	initial := map[cpu.Register]int{}
	value := 0

	for reg := range registers {
		initial[reg] = value
		value++
	}

	expected := map[cpu.Register]int{}
	maps.Copy(expected, initial)

	for _, move := range moves {
		expected[move.Destination] = initial[move.Source]
	}

	actual := map[cpu.Register]int{}
	maps.Copy(actual, initial)

	for _, move := range scheduled {
		actual[move.Destination] = actual[move.Source]
	}

	for reg := range registers {
		assert.Equal(t, actual[reg], expected[reg])
	}
}
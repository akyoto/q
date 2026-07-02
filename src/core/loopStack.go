package core

// loopStack remembers all active loops during compilation.
type loopStack struct {
	loops []*Loop
}

// Current returns the innermost loop or nil if there is none.
func (s *loopStack) Current() *Loop {
	if len(s.loops) == 0 {
		return nil
	}

	return s.loops[len(s.loops)-1]
}

// Pop removes an element from the stack.
func (s *loopStack) Pop() {
	s.loops = s.loops[:len(s.loops)-1]
}

// Push pushes a new element to the stack.
func (s *loopStack) Push(loop *Loop) {
	s.loops = append(s.loops, loop)
}
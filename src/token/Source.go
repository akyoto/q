package token

// Source is a region that has a start and end position in the source file.
type Source struct {
	start Position
	end   Position
}

// NewSource creates a new source region.
func NewSource(start Position, end Position) Source {
	return Source{start: start, end: end}
}

// End returns the end position.
func (s Source) End() Position {
	return s.end
}

// SetEnd sets the end position.
func (s *Source) SetEnd(end Position) {
	s.end = end
}

// SetStart sets the start position.
func (s *Source) SetStart(start Position) {
	s.start = start
}

// Start returns the start position.
func (s Source) Start() Position {
	return s.start
}

// StringFrom returns the source code.
func (s Source) StringFrom(code []byte) string {
	return string(code[s.Start():s.End()])
}
package token

// Source is a region that has a start and end position in the source file.
type Source struct {
	StartPos Position
	EndPos   Position
}

// Start returns the start position.
func (s Source) Start() Position {
	return s.StartPos
}

// End returns the end position.
func (s Source) End() Position {
	return s.EndPos
}

// StringFrom returns the source code.
func (s Source) StringFrom(code []byte) string {
	return string(code[s.Start():s.End()])
}
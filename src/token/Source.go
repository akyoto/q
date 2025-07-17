package token

// Source is a region that has a start and end position in the source file.
type Source struct {
	StartPos Position
	EndPos   Position
}
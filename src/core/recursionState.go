package core

// recursionState represents the processing state while parsing struct types.
// It is used to detect cycles when a struct references another struct that is
// in "Started" state.
type recursionState byte

const (
	notStarted recursionState = iota
	started
	finished
)
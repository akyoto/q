package exe

// Discard implements a no-op WriteSeeker.
type Discard struct{}

func (w *Discard) Write(_ []byte) (int, error)        { return 0, nil }
func (w *Discard) Seek(_ int64, _ int) (int64, error) { return 0, nil }
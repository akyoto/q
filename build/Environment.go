package build

import "sync"

// Environment represents the global state.
type Environment struct {
	functions sync.Map
}

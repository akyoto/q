package build

import (
	"sync"
)

var (
	// Pool for function calls
	functionCallPool = sync.Pool{
		New: func() interface{} {
			return &FunctionCall{}
		},
	}
)

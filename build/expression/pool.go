package expression

import "sync"

var pool = sync.Pool{
	New: func() interface{} {
		return &Expression{}
	},
}

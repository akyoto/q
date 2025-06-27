package set_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/go/assert"
)

func TestOrdered(t *testing.T) {
	s := set.Ordered[int]{}
	assert.Equal(t, s.Count(), 0)
	s.Add(1)
	assert.Equal(t, s.Count(), 1)
	s.Add(2)
	assert.Equal(t, s.Count(), 2)
	s.Add(3)
	assert.Equal(t, s.Count(), 3)
	s.Add(1)
	assert.Equal(t, s.Count(), 3)

	for element := range s.All() {
		if element == 2 {
			break
		}
	}
}
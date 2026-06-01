package set_test

import (
	"testing"

	"git.urbach.dev/cli/q/src/set"
	"git.urbach.dev/go/assert"
)

func TestBringToBack(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 3, 4, 5, 2}
	set.BringToBack(a, 1)
	assert.DeepEqual(t, a, b)
}

func TestBringToBackPart(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 3, 4, 2, 5}
	set.BringToBack(a[1:4], 0)
	assert.DeepEqual(t, a, b)
}

func TestBringToBackInvalid(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	set.BringToBack(a, -1)
	set.BringToFront(a, 5)
	assert.DeepEqual(t, a, b)
}

func TestBringToFront(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 1, 2, 3, 5}
	set.BringToFront(a, 3)
	assert.DeepEqual(t, a, b)
}

func TestBringToFrontPart(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 4, 2, 3, 5}
	set.BringToFront(a[1:4], 2)
	assert.DeepEqual(t, a, b)
}

func TestBringToFrontInvalid(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	set.BringToFront(a, -1)
	set.BringToFront(a, 5)
	assert.DeepEqual(t, a, b)
}
package core

import (
	"sync"

	"git.urbach.dev/cli/q/src/types"
)

// TypeCache contains reusable type objects.
type TypeCache struct {
	pointerTypes  map[types.Type]types.Type
	resourceTypes map[types.Type]types.Type
	sliceTypes    map[types.Type]types.Type
	pointerMutex  sync.Mutex
	resourceMutex sync.Mutex
	sliceMutex    sync.Mutex
}

// Pointer returns the type that points to the given type.
func (c *TypeCache) Pointer(typ types.Type) types.Type {
	c.pointerMutex.Lock()
	defer c.pointerMutex.Unlock()
	existing, exists := c.pointerTypes[typ]

	if !exists {
		existing = &types.Pointer{To: typ}
		c.pointerTypes[typ] = existing
	}

	return existing
}

// Resource returns the type that is a resource of the given type.
func (c *TypeCache) Resource(typ types.Type) types.Type {
	c.resourceMutex.Lock()
	defer c.resourceMutex.Unlock()
	existing, exists := c.resourceTypes[typ]

	if !exists {
		existing = &types.Resource{Of: typ}
		c.resourceTypes[typ] = existing
	}

	return existing
}

// Slice returns the type that is a slice of the given type.
func (c *TypeCache) Slice(typ types.Type) types.Type {
	c.sliceMutex.Lock()
	defer c.sliceMutex.Unlock()
	existing, exists := c.sliceTypes[typ]

	if !exists {
		existing = types.Slice(typ, "[]"+typ.Name())
		c.sliceTypes[typ] = existing
	}

	return existing
}
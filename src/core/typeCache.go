package core

import (
	"sync"

	"git.urbach.dev/cli/q/src/types"
)

// typeCache contains reusable type objects.
type typeCache struct {
	pointerTypes  sync.Map
	resourceTypes sync.Map
	sliceTypes    sync.Map
}

// Pointer returns the type that points to the given type.
func (c *typeCache) Pointer(typ types.Type) types.Type {
	existing, ok := c.pointerTypes.Load(typ)

	if ok {
		return existing.(types.Type)
	}

	new := &types.Pointer{To: typ}
	existing, _ = c.pointerTypes.LoadOrStore(typ, new)
	return existing.(types.Type)
}

// Resource returns the type that is a resource of the given type.
func (c *typeCache) Resource(typ types.Type) types.Type {
	existing, ok := c.resourceTypes.Load(typ)

	if ok {
		return existing.(types.Type)
	}

	new := &types.Resource{Of: typ}
	existing, _ = c.resourceTypes.LoadOrStore(typ, new)
	return existing.(types.Type)
}

// Slice returns the type that is a slice of the given type.
func (c *typeCache) Slice(typ types.Type) types.Type {
	existing, ok := c.sliceTypes.Load(typ)

	if ok {
		return existing.(types.Type)
	}

	new := types.Slice(typ, "[]"+typ.Name())
	existing, _ = c.sliceTypes.LoadOrStore(typ, new)
	return existing.(types.Type)
}
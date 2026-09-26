package set

import (
	"maps"
)

// CowMap is a copy-on-write map sharing storage with another map until its first write.
type CowMap[K comparable, V comparable] struct {
	raw    map[K]V
	shared bool
}

// Count returns the number of entries.
func (m *CowMap[K, V]) Count() int {
	return len(m.raw)
}

// Get looks up a key.
func (m *CowMap[K, V]) Get(key K) (V, bool) {
	value, exists := m.raw[key]
	return value, exists
}

// Raw returns the underlying map for reading purposes.
func (m *CowMap[K, V]) Raw() map[K]V {
	return m.raw
}

// RemoveByValue deletes the first key with the given value.
func (m *CowMap[K, V]) RemoveByValue(value V) {
	for key, existing := range m.raw {
		if existing == value {
			if m.shared {
				m.raw = maps.Clone(m.raw)
				m.shared = false
			}

			delete(m.raw, key)
			return
		}
	}
}

// Set adds or changes a key, copying the map if it's still shared.
func (m *CowMap[K, V]) Set(key K, value V) {
	if m.shared {
		m.raw = maps.Clone(m.raw)
		m.shared = false
	}

	if m.raw == nil {
		m.raw = make(map[K]V)
	}

	m.raw[key] = value
}

// Share assigns the initial map that shares storage with this one.
func (m *CowMap[K, V]) Share(values map[K]V) {
	m.raw = values
	m.shared = true
}
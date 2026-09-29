package data

// SetImmutable sets the readonly data for the given label.
func (data *Data) SetImmutable(label string, bytes []byte) {
	if data.Immutable == nil {
		data.Immutable = map[string][]byte{}
	}

	data.Immutable[label] = bytes
}

// SetMutable sets the writable data for the given label.
func (data *Data) SetMutable(label string, bytes []byte) {
	if data.Mutable == nil {
		data.Mutable = map[string][]byte{}
	}

	data.Mutable[label] = bytes
}

// SetTLS sets the thread-local data for the given label.
func (data *Data) SetTLS(label string, bytes []byte) {
	if data.TLS == nil {
		data.TLS = map[string][]byte{}
	}

	data.TLS[label] = bytes
}
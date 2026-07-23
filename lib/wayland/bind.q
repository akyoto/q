import io

bind(state *State, buffer string, request *BindRequest) -> error {
	padded := pad(request.interface.len + 1)
	size := bindSize(padded)

	if size > buffer.len {
		return -1
	}

	w := newWriter(buffer.ptr, buffer.len)
	w = write32(w, state.wl_registry)
	w = write16(w, wl_registry_bind)
	w = write16(w, size)
	w = write32(w, request.name)
	w = writeString(w, request.interface)
	w = write32(w, request.version)
	w = write32(w, request.id)
	_, err := io.writeTo(state.socket, buffer[..w.len])
	return err
}

bindSize(padded uint) -> uint16 {
	return header_size + 4 + 4 + padded + 4 + 4
}
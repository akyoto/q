import io
import net
import strings
import wayland

main() {
	socket, err := wayland.connect()

	if err != 0 {
		io.write("error: ")
		io.write(err)
		return
	}

	state := new(wayland.State) {
		id: wayland.wl_display_id,
		socket: socket
	}

	buffer := new(byte, 4096)
	err := communicate(state, buffer)

	if err != 0 {
		io.write("error: ")
		io.write(err)
	}

	deleteShm(state)
	delete(state)
	delete(buffer)
	net.close(socket)
}

communicate(state *wayland.State, buffer string) -> error {
	err := getRegistry(state, buffer)

	if err != 0 {
		return err
	}

	err := read(state, buffer)

	if err != 0 {
		return err
	}

	if state.wl_compositor_name == 0 || state.wl_shm_name == 0 || state.xdg_wm_base_name == 0 {
		return -1
	}

	state.wl_compositor = wayland.newId(state)
	err := bindCompositor(state, buffer, "wl_compositor")

	if err != 0 {
		return err
	}

	state.wl_shm = wayland.newId(state)
	err := bindShm(state, buffer, "wl_shm")

	if err != 0 {
		return err
	}

	state.xdg_wm_base = wayland.newId(state)
	err := bindXdgWmBase(state, buffer, "xdg_wm_base")

	if err != 0 {
		return err
	}

	err := read(state, buffer)

	if err != 0 {
		return err
	}

	err := createShm(state)

	if err != 0 {
		return err
	}

	return 0
}

read(state *wayland.State, buffer string) -> error {
	n, err := io.readFrom(state.socket, buffer)

	if err != 0 {
		return err
	}

	if n == 0 {
		return 0
	}

	pos := 0

	loop {
		if pos + wayland.header_size >= n {
			return 0
		}

		pos += handleMessage(state, buffer[pos..n])
	}
}

handleMessage(state *wayland.State, msg string) -> int {
	header := msg.ptr as *wayland.Header
	//assert header.size <= msg.len

	if header.id == state.wl_registry {
		name := [msg.ptr + wayland.header_size as *uint32]
		len := [msg.ptr + wayland.header_size + 4 as *uint32]
		start := wayland.header_size + 8
		end := start + len
		interface := msg[start..end-1]
		version := [msg.ptr + end as *uint32]
		handleGlobal(state, name, interface, version)
	}

	if header.id == wayland.wl_display_id && header.opcode == wayland.wl_display_error {
		io.writeLine("wl_display::error")
	}

	return header.size as int
}

handleGlobal(state *wayland.State, name uint32, interface string, _version uint32) {
	io.write("[")
	io.write(name as int)
	io.write("] ")
	io.write(interface)
	io.write("\n")

	switch {
		strings.equal(interface, "wl_compositor") {
			state.wl_compositor_name = name
		}
		strings.equal(interface, "wl_shm") {
			state.wl_shm_name = name
		}
		strings.equal(interface, "xdg_wm_base") {
			state.xdg_wm_base_name = name
		}
	}
}

getRegistry(state *wayland.State, buffer string) -> error {
	state.wl_registry = wayland.newId(state)
	size := wayland.header_size + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, wayland.wl_display_id)
	ptr = wayland.write16(ptr, wayland.wl_display_get_registry)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_registry)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

createPool(state *wayland.State, buffer string) -> error {
	state.wl_shm_pool = wayland.newId(state)
	size := wayland.header_size + 4 + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_shm)
	ptr = wayland.write16(ptr, wayland.wl_shm_create_pool)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_shm_pool)
	ptr = wayland.write32(ptr, state.shm_size)
	// TODO: Send fd as ancillary data
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindCompositor(state *wayland.State, buffer string, interface string) -> error {
	padded := wayland.pad(interface.len + 1)
	size := bindSize(padded)

	if size > buffer.len {
		return -1
	}

	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.wl_registry_bind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_compositor_name)
	ptr = wayland.writeString(ptr, interface)
	ptr = wayland.write32(ptr, 4)
	ptr = wayland.write32(ptr, state.wl_compositor)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindShm(state *wayland.State, buffer string, interface string) -> error {
	padded := wayland.pad(interface.len + 1)
	size := bindSize(padded)

	if size > buffer.len {
		return -1
	}

	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.wl_registry_bind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_shm_name)
	ptr = wayland.writeString(ptr, interface)
	ptr = wayland.write32(ptr, 1)
	ptr = wayland.write32(ptr, state.wl_shm)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindXdgWmBase(state *wayland.State, buffer string, interface string) -> error {
	padded := wayland.pad(interface.len + 1)
	size := bindSize(padded)

	if size > buffer.len {
		return -1
	}

	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.wl_registry_bind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.xdg_wm_base_name)
	ptr = wayland.writeString(ptr, interface)
	ptr = wayland.write32(ptr, 2)
	ptr = wayland.write32(ptr, state.xdg_wm_base)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindSize(padded uint) -> uint16 {
	return wayland.header_size + 4 + 4 + padded + 4 + 4
}
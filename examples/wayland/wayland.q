import io
import mem
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
		id: 1,
		socket: socket
	}

	buffer := new(byte, 4096)
	err := communicate(state, buffer)

	if err != 0 {
		io.write("error: ")
		io.write(err)
	}

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

	mem.zero(buffer)
	size := 0 as uint
	state.wl_compositor = wayland.newId(state)
	err := bindCompositor(state, buffer[size..], state.wl_compositor)

	if err != 0 {
		return err
	}

	state.wl_shm = wayland.newId(state)
	err := bindShm(state, buffer[size..], state.wl_shm)

	if err != 0 {
		return err
	}

	state.xdg_wm_base = wayland.newId(state)
	err := bindXdgWmBase(state, buffer[size..], state.xdg_wm_base)

	if err != 0 {
		return err
	}

	err := read(state, buffer)

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
		if pos + wayland.headerSize >= n {
			return 0
		}

		pos += handleMessage(state, buffer[pos..n])
	}
}

handleMessage(state *wayland.State, msg string) -> int {
	header := msg.ptr as *wayland.Header
	//assert header.size <= msg.len

	if header.id == state.wl_registry {
		name := [msg.ptr + wayland.headerSize as *uint32]
		len := [msg.ptr + wayland.headerSize + 4 as *uint32]
		start := wayland.headerSize + 8
		end := start + len
		interface := msg[start..end-1]
		version := [msg.ptr + end as *uint32]
		handleGlobal(state, name, interface, version)
	}

	if header.id == wayland.displayId && header.opcode == wayland.displayError {
		//io.writeLine("wl_display::error")
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
	size := wayland.headerSize + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, wayland.displayId)
	ptr = wayland.write16(ptr, wayland.displayGetRegistry)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_registry)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindCompositor(state *wayland.State, buffer string, id uint32) -> error {
	padded := buffer.len + 1
	padded = (padded + 3) & -4
	size := wayland.headerSize + 4 + 4 + padded + 4 + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.registryBind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_compositor_name)
	ptr = wayland.writeString(ptr, "wl_compositor")
	ptr = wayland.write32(ptr, 4)
	ptr = wayland.write32(ptr, id)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindShm(state *wayland.State, buffer string, id uint32) -> error {
	padded := buffer.len + 1
	padded = (padded + 3) & -4
	size := wayland.headerSize + 4 + 4 + padded + 4 + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.registryBind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.wl_shm_name)
	ptr = wayland.writeString(ptr, "wl_shm")
	ptr = wayland.write32(ptr, 1)
	ptr = wayland.write32(ptr, id)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

bindXdgWmBase(state *wayland.State, buffer string, id uint32) -> error {
	padded := buffer.len + 1
	padded = (padded + 3) & -4
	size := wayland.headerSize + 4 + 4 + padded + 4 + 4
	ptr := buffer.ptr
	ptr = wayland.write32(ptr, state.wl_registry)
	ptr = wayland.write16(ptr, wayland.registryBind)
	ptr = wayland.write16(ptr, size)
	ptr = wayland.write32(ptr, state.xdg_wm_base_name)
	ptr = wayland.writeString(ptr, "xdg_wm_base")
	ptr = wayland.write32(ptr, 2)
	ptr = wayland.write32(ptr, id)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}
import io
import mem
import net
import strings
import wayland

main() {
	socket, err := wayland.connect()

	if err != 0 {
		io.write("connect error: ")
		io.write(err)
		return
	}

	state := new(wayland.State) {
		id: 1,
		socket: socket
	}

	buffer := new(byte, 4096)
	err := getRegistry(state, buffer)

	if err != 0 {
		io.write("send error: ")
		io.write(err)
		delete(state)
		delete(buffer)
		net.close(socket)
		return
	}

	loop {
		err := readMessage(state, buffer)

		if err != 0 {
			delete(state)
			delete(buffer)
			net.close(socket)
			return
		}
	}
}

getRegistry(state *wayland.State, buffer string) -> error {
	state.wl_registry = wayland.newId(state)
	size := wayland.headerSize + 4
	wayland.write32(buffer, wayland.displayId)
	wayland.write16(buffer[4..], wayland.displayGetRegistry)
	wayland.write16(buffer[6..], size)
	wayland.write32(buffer[8..], state.wl_registry)
	_, err := io.writeTo(state.socket, buffer[..size])
	return err
}

readMessage(state *wayland.State, buffer string) -> error {
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

	return header.size as int
}

handleGlobal(state *wayland.State, name uint32, interface string, version uint32) {
	io.write("[")
	io.write(name as int)
	io.write("] ")
	io.write(interface)
	io.write("\n")

	switch {
		strings.equal(interface, "wl_compositor") {
			state.wl_compositor = bind(state, name, interface, version)
		}
		strings.equal(interface, "wl_shm") {
			state.wl_shm = bind(state, name, interface, version)
		}
		strings.equal(interface, "xdg_wm_base") {
			state.xdg_wm_base = bind(state, name, interface, version)
		}
	}
}

bind(state *wayland.State, name uint32, interface string, _version uint32) -> uint32 {
	size := wayland.headerSize + 4 + 4 + interface.len + 4 + 4
	buffer := new(byte, size)
	wayland.write32(buffer, state.wl_registry)
	wayland.write16(buffer[4..], wayland.registryBind)
	wayland.write16(buffer[6..], size)
	wayland.write32(buffer[8..], name)
	wayland.write32(buffer[12..], interface.len as uint32)
	mem.copy(buffer[16..], interface)
	// wayland.write32(buffer[16+interface.len..], version)
	id := wayland.newId(state)
	wayland.write32(buffer[16+interface.len+4..], id)
	//io.writeTo(state.socket, buffer[..size])
	delete(buffer)
	return id
}
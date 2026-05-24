import io
import net
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
	state.registry = wayland.newId(state)
	size := wayland.headerSize + 4
	wayland.write32(buffer, wayland.displayId)
	wayland.write16(buffer[4..], wayland.displayGetRegistry)
	wayland.write16(buffer[6..], size)
	wayland.write32(buffer[8..], state.registry)
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

	if header.id == state.registry {
		io.write("from:")
		io.write(header.id as int)
		io.write(" opcode:")
		io.write(header.opcode as int)
		io.write(" size:")
		io.write(header.size as int)
		io.write(" contents:")
		id := [msg.ptr + wayland.headerSize as *uint32]
		io.write(" id:")
		io.write(id as int)
		io.write(" name:")
		len := [msg.ptr + wayland.headerSize + 4 as *uint32]
		start := wayland.headerSize + 8
		end := start + len
		contents := msg[start..end]
		io.writeLine(contents)
	}

	return header.size as int
}
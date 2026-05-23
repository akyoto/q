import io
import mem
import net
import wayland

main() {
	io.writeLine("connecting...")
	socket, err := wayland.connect()

	if err != 0 {
		io.write(err)
		return
	}

	io.writeLine("connected.")
	io.writeLine("sending message...")
	buffer := mem.alloc(4096)
	size := wayland.header_size + 4
	wayland.write32(buffer, wayland.display_object_id)
	wayland.write16(buffer[4..], wayland.wl_display_get_registry_opcode)
	wayland.write16(buffer[6..], size)
	wayland.write32(buffer[8..], 2)
	io.writeTo(socket, buffer[..size])
	io.writeLine("sent.")
	io.writeLine("receiving message...")
	n, err := io.readFrom(socket, buffer)

	if err != 0 {
		io.write(err)
		mem.free(buffer)
		net.close(socket)
		return
	}

	io.writeLine("received.")
	io.writeLine("size:")
	io.writeLine(n)
	io.writeLine("msg:")
	io.writeLine(buffer[wayland.header_size..n])
	mem.free(buffer)
	net.close(socket)
}
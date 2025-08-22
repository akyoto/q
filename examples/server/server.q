import io
import net

main() {
	socket := net.listenTCP(0, 8080)

	if socket < 0 {
		io.write("socket error\n")
		return
	}

	io.write("http://127.0.0.1:8080\n")
	io.write("listening...\n")

	loop {
		conn := net.accept(socket)

		if conn >= 0 {
			net.send(conn, "HTTP/1.0 200 OK\r\nContent-Length: 6\r\n\r\nHello\n")
			net.close(conn)
		} else {
			io.write("accept error\n")
		}
	}
}
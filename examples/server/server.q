import io
import net
import os

main() {
	socket := net.socket(2, 1, 0)

	if socket < 0 {
		io.write("socket error\n")
		os.exit(1)
	}

	if net.bind(socket, 8080) != 0 {
		io.write("bind error\n")
		os.exit(1)
	}

	if net.listen(socket, 128) != 0 {
		io.write("listen error\n")
		os.exit(1)
	}

	io.write("http://127.0.0.1:8080\n")
	io.write("listening...\n")

	loop {
		conn := net.accept(socket, 0, 0)

		if conn >= 0 {
			net.send(conn, "HTTP/1.0 200 OK\r\nContent-Length: 6\r\n\r\nHello\n")
			net.close(conn)
		} else {
			io.write("accept error\n")
		}
	}
}
import io
import net
import tcp

main() {
	socket, err := tcp.listen(8080)

	if err != 0 {
		io.write("listen error\n")
		io.write(err)
		return
	}

	io.write("http://[::1]:8080\n")
	io.write("listening...\n")

	loop {
		conn, err := net.accept(socket)

		if err == 0 {
			net.send(conn, "HTTP/1.0 200 OK\r\nContent-Length: 6\r\n\r\nHello\n")
			net.close(conn)
		} else {
			io.write("accept error\n")
		}
	}
}
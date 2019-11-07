main() {
	out(11)
}

out(msgLength) {
	id = 1
	fd = 1
	msg = "Hello World"

	syscall(id, fd, msg, msgLength)
}

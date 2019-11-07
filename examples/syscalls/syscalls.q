main() {
	out()
}

out() {
	id = 1
	fd = 1
	msg = "Hello World"
	msgLength = 11

	syscall(id, fd, msg, msgLength)
}

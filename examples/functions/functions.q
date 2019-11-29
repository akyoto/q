main() {
	a = sum(1, 2)
	b = sum(3, 4)
	c = sum(a, b)
	d = write(1, "1234567890", c)
	exit(d)
}

sum(a, b) {
	return a + b
}

write(fd, msg, length) {
	return syscall(1, fd, msg, length)
}

exit(code) {
	syscall(60, code)
}

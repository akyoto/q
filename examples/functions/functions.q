main() {
	a = sum(1, 2)
	b = sum(3, 4)
	c = sum(a, b)
	write(1, "123456789\n", c)
}

sum(a, b) {
	return a + b
}

write(fd, msg, length) {
	return syscall(1, fd, msg, length)
}

main() {
	stdout = 1
	a = sum(1, 2)
	b = sum(3, 4)
	c = sum(a, b)
	d = write(stdout, "1234567890", c)
	exit(d)
}

sum(a, b) {
	return a + b
}

write(fd, msg, length) {
	tmp1 = fd
	tmp2 = msg
	tmp3 = length
	return syscall(1, tmp1, tmp2, tmp3)
}

exit(code) {
	tmp1 = code
	syscall(60, tmp1)
}

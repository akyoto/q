main() {
	a = sum(1, 2)
	b = sum(3, 4)
	c = sum(a, b)
	d = write("1234567890", c)
	exit(d)
}

sum(a, b) {
	return a + b
}

write(msg, length) {
	tmp1 = msg
	tmp2 = length
	return syscall(1, 1, tmp1, tmp2)
}

exit(code) {
	tmp1 = code
	syscall(60, tmp1)
}

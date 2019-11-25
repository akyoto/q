main() {
	a = sum(1, 2)
	b = sum(3, 4)
	c = sum(a, b)
	syscall(60, c)
}

sum(a, b) {
	return a + b
}

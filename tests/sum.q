main() {
	t1 := sum(1, 2)
	t2 := sum(3, 4)
	t3 := sum(t1, t2)
	syscall(60, t3)
}

sum(a int, b int) -> int {
	return a + b
}
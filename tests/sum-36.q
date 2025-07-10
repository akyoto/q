import os

main() {
	t1 := sum(1, 2)
	t2 := sum(3, 4)
	t3 := sum(t1, t2)
	t4 := sum(5, 6)
	t5 := sum(7, 8)
	t6 := sum(t4, t5)
	t7 := sum(t3, t6)
	os.exit(t7)
}

sum(a int, b int) -> int {
	return a + b
}
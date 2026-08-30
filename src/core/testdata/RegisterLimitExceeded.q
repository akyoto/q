main() {
	assert fn(1, 2, 3, 4, 5, 6, 7, 8) == 1
}

fn(a int, _b int, _c int, _d int, _e int, _f int, _g int, _h int) -> int {
	return a
}
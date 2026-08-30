Small {
	a int8
	b int8
	c int8
	d int8
}

Big {
	a int64
	b int64
	c int64
	d int64
}

main() {
	assert f(Small{a: 1, b: 2, c: 3, d: 4}, 42) == 42
	assert g(Big{a: 1, b: 2, c: 3, d: 4}, 42) == 42
}

f(_p Small, x int) -> int {
	return x
}

g(_p Big, x int) -> int {
	return x
}
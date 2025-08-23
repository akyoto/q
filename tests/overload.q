main() {
	assert f(123) == 1
	assert f("Hello") == 2
}

f(_ int) -> int {
	return 1
}

f(_ string) -> int {
	return 2
}
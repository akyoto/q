main() {
	assert f(123) == 1
	assert f(123, 456) == 2
	assert f(123, 456, 789) == 3
	assert f("Hello") == 4
}

f(_ int) -> int {
	return 1
}

f(_ int, _ int) -> int {
	return 2
}

f(_ int, _ int, _ int) -> int {
	return 3
}

f(_ string) -> int {
	return 4
}
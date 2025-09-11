main() {
	x := new(*int)
	[x] = new(int)
	[[x]] = 42
	assert [[x]] == 42
}
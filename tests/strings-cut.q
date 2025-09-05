import strings

main() {
	hello, world := strings.cut("Hello World", " ")
	assert strings.equal(hello, "Hello") == true
	assert strings.equal(world, "World") == true

	sekai, empty := strings.cut("世界", ":")
	assert strings.equal(sekai, "世界") == true
	assert strings.equal(empty, "") == true
}
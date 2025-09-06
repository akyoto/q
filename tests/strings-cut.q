import strings

main() {
	hello, world := strings.cut("Hello World", " ")
	assert strings.equal(hello, "Hello")
	assert strings.equal(world, "World")

	sekai, empty := strings.cut("世界", ":")
	assert strings.equal(sekai, "世界")
	assert strings.equal(empty, "")
}
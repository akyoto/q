import strings

main() {
	hello, world, err := strings.cut("Hello World", " ")
	assert err == 0
	assert hello == "Hello"
	assert world == "World"

	_, _, err := strings.cut("世界", ":")
	assert err != 0
}
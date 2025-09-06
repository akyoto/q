import strings

main() {
	assert strings.hasPrefix("Hello", "")
	assert strings.hasPrefix("Hello", "H")
	assert strings.hasPrefix("Hello", "He")
	assert strings.hasPrefix("Hello", "World") == false
	assert strings.hasPrefix("Hello", "Hello World") == false
}
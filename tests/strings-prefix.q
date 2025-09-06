import strings

main() {
	assert strings.hasPrefix("Hello", "") == true
	assert strings.hasPrefix("Hello", "H") == true
	assert strings.hasPrefix("Hello", "He") == true
	assert strings.hasPrefix("Hello", "World") == false
	assert strings.hasPrefix("Hello", "Hello World") == false
}
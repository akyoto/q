import strings

main() {
	assert strings.hasSuffix("Hello", "") == true
	assert strings.hasSuffix("Hello", "o") == true
	assert strings.hasSuffix("Hello", "lo") == true
	assert strings.hasSuffix("Hello", "World") == false
	assert strings.hasSuffix("Hello", "Hello World") == false
}
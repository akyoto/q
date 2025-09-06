import strings

main() {
	assert strings.hasSuffix("Hello", "")
	assert strings.hasSuffix("Hello", "o")
	assert strings.hasSuffix("Hello", "lo")
	assert !strings.hasSuffix("Hello", "World")
	assert !strings.hasSuffix("Hello", "Hello World")
}
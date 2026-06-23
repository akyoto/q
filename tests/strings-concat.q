import strings

main() {
	combined := strings.concat("Hello", "World")
	assert combined == "HelloWorld"
	assert combined != "WorldHello"
}
import strings

main() {
	combined := "Hello" ++ "World"
	assert strings.equal(combined, "HelloWorld")
	assert !strings.equal(combined, "WorldHello")
}
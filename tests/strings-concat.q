import mem
import strings

main() {
	combined := strings.concat("Hello", "World")
	assert strings.equal(combined, "HelloWorld") == true
	mem.free(combined)
}
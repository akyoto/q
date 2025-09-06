import mem
import strings

main() {
	combined := strings.concat("Hello", "World")
	assert strings.equal(combined, "HelloWorld")
	mem.free(combined)
}
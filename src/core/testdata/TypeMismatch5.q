import mem

main() {
	buffer := mem.alloc(1)
	delete(buffer)
}
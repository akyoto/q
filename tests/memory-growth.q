import mem

main() {
	_ := new(byte, 1)
	b := new(byte, 20000000)
	assert mem.heap.max - b.ptr >= b.len
}
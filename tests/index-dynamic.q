import mem

main() {
	buffer := mem.alloc(4)
	a := buffer.ptr

	loop i := 0..4 {
		a[i] = i
		assert a[i] == i
	}
}
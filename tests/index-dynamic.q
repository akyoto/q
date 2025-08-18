import mem

main() {
	a := mem.alloc(4)

	loop i := 0..4 {
		a[i] = i
		assert a[i] == i
	}
}
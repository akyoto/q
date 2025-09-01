main() {
	a := "aa"
	b := "ab"

	loop i := 0..a.len {
		if a[i] != b[i] {
			return
		}
	}
}
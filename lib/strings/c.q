import mem

c(s string) -> !string {
	cstr := mem.alloc(s.len + 1)

	loop i := 0..s.len {
		cstr[i] = s[i]
	}

	return cstr
}
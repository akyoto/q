import mem

string(s string) -> !string {
	cstr := mem.alloc(s.len + 1)
	mem.copy(cstr, s)
	return cstr
}
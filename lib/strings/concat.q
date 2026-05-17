import mem

concat(a string, b string) -> !string {
	combined := mem.alloc(a.len + b.len)
	mem.copy(combined, a)
	mem.copy(combined[a.len..], b)
	return combined
}
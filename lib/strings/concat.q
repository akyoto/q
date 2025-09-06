import mem

concat(a string, b string) -> !string {
	combined := mem.alloc(a.len + b.len)

	loop i := 0..a.len {
		combined[i] = a[i]
	}

	loop i := 0..b.len {
		combined[a.len+i] = b[i]
	}

	return combined
}
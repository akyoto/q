import mem

Writer {
	ptr *byte
	len uint
	cap uint
}

newWriter(ptr *byte, cap uint) -> Writer {
	return Writer{ptr: ptr, len: 0, cap: cap}
}

write16(w Writer, v uint16) -> Writer {
	[w.ptr + w.len as *uint16] = v
	w.len += 2
	return w
}

write32(w Writer, v uint32) -> Writer {
	[w.ptr + w.len as *uint32] = v
	w.len += 4
	return w
}

writeString(w Writer, s string) -> Writer {
	length := s.len + 1
	w = write32(w, length as uint32)
	padded := pad(length)
	mem.copy(string{ptr: w.ptr + w.len, len: s.len}, s)
	mem.zero(string{ptr: w.ptr + w.len + s.len, len: padded - s.len})
	w.len += padded
	return w
}
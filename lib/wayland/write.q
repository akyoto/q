import mem

write16(ptr *byte, v uint16) -> *byte {
	[ptr as *uint16] = v
	return ptr + 2
}

write32(ptr *byte, v uint32) -> *byte {
	[ptr as *uint32] = v
	return ptr + 4
}

writeString(ptr *byte, s string) -> *byte {
	len := s.len + 1
	ptr = write32(ptr, len as uint32)
	padded := (len + 3) & -4

	loop i := 0..s.len {
		ptr[i] = s.ptr[i]
	}

	mem.zero(string{ptr: ptr+s.len, len: padded-s.len})
	return ptr + padded
}
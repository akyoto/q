fromInt(x int, base int, buffer []byte) -> []byte {
	negative := x < 0
	end := buffer.ptr + buffer.len
	tmp := end

	loop {
		digit := x % base
		x /= base
		tmp -= 1
		[tmp] = "FEDCBA9876543210123456789ABCDEF"[digit+15]

		if x == 0 {
			if negative {
				tmp -= 1
				[tmp] = '-'
			}

			return []byte{
				ptr: tmp,
				len: end - tmp as uint,
			}
		}
	}
}

fromInt(x uint, base uint, buffer []byte) -> []byte {
	end := buffer.ptr + buffer.len
	tmp := end

	loop {
		digit := x % base
		x /= base
		tmp -= 1
		[tmp] = "0123456789ABCDEF"[digit]

		if x == 0 {
			return []byte{
				ptr: tmp,
				len: end - tmp as uint,
			}
		}
	}
}
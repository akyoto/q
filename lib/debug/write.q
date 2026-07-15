import io

write(bytes []byte) {
	loop i := 0..bytes.len {
		n := bytes[i]
		writeDigit(n / 16 % 16 as uint)
		writeDigit(n % 16 as uint)

		if i % 8 == 7 {
			io.write("\n")
			loop.next()
		}

		if i + 1 < bytes.len {
			io.write(" ")
		}
	}
}

writeDigit(n uint) {
	write("0123456789ABCDEF"[n..n+1])
}
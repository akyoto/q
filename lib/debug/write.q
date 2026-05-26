import io

write(bytes []byte) {
	loop i := 0..bytes.len {
		n := bytes[i]
		io.writeDigit(n / 16 % 16 as int)
		io.writeDigit(n % 16 as int)

		if i % 8 == 7 {
			io.write("\n")
			loop.next()
		}

		if i + 1 < bytes.len {
			io.write(" ")
		}
	}
}
import io

main() {
	data := "abcdefghijklmn"
	state := 2147483648 as uint
	output := new(byte, data.len * 2)
	outputPos := 0 as uint

	loop i := 0..data.len {
		c := data[data.len - i - 1]
		state, outputPos = encode(state, output, outputPos, c)
	}

	decoded := new(byte, data.len)
	d := 0 as byte

	loop i := 0..data.len {
		state, d, outputPos = decode(state, output, outputPos)
		decoded[i] = d
	}

	io.writeLine(d as int)
	io.writeLine(decoded.len)
	io.writeLine(outputPos)
}

encode(state uint, output []byte, outputPos uint, _value byte) -> (uint, uint) {
	xMax := 134217728 as uint
	xMax *= 100

	loop {
		if state < xMax {
			loop.stop()
		}

		output[outputPos] = (state & 255) as byte
		outputPos += 1
		state >>= 8
	}

	return state, outputPos
}

decode(state uint, input []byte, inputPos uint) -> (uint, byte, uint) {
	symbol := 0 as byte

	loop {
		if state >= 2147483648 {
			loop.stop()
		}

		inputPos -= 1

		x := input[inputPos] as uint
		state = (state << 8) + x
	}

	return state, symbol, inputPos
}
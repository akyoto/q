import io

main() {
	output := new(byte, 28)
	state := 2147483648 as uint
	outputPos := 0 as uint
	decoded := new(byte, 14)
	d := 0 as byte

	loop i := 0..3 {
		state, d, outputPos = decode(state, output, outputPos)
		decoded[i] = d
	}

	io.writeLine(d as int)
	io.writeLine(decoded.len)
	io.writeLine(outputPos)
}

decode(state uint, _input []byte, input_pos uint) -> (uint, byte, uint) {
	symbol := 0 as byte

	loop {
		if state >= 2147483648 {
			loop.stop()
		}

		input_pos -= 1
		state += 1
	}

	return state, symbol, input_pos
}
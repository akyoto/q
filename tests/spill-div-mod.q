import io

const {
	TOTAL = 4096
}

global {
	freq []int
	bound int
}

encode(input []byte, output []byte) -> int {
	state := 65536
	pos := 0

	loop i := 0..bound {
		s := input[i]
		f := freq[s]
		c := input[i + 1]

		loop {
			if state < f * TOTAL {
				loop.stop()
			}

			output[pos] = (state & 255) as byte
			pos += 1
			state >>= 8
		}

		io.writeLine(state)
		state = (state / f) * TOTAL + c + (state % f)
	}

	return pos
}

main() {
	freq = new(int, 256)

	loop i := 0..256 {
		freq[i] = i % 7 + 1
	}

	bound = 10
	input := new(byte, bound + 1)

	loop i := 0..bound {
		input[i] = i as byte
	}

	encoded := new(byte, 64)
	n := encode(input, encoded)
	assert n == 13
}
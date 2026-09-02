import io
import run

main() {
	x := 3
	y := 7

	if x == 0 {
		io.write("A")
	}

	if x == 0 {
		io.write("B")
	} else {
		io.write("C")
	}

	if x != 0 || y == 0 {
		io.write("D")
	}

	if x < 0 && y < 0 {
		io.write("E")
	} else {
		io.write("F")
	}

	loop 0..2 {
		if x == 0 {
			io.write("G")
		}
	}

	if !(x == 0) {
		io.write("H")
	} else {
		io.write("I")
	}

	if x == 0 {
		io.write("J")
	}

	assert x != 0

	if y == 0 {
		run.crash()
	}

	io.write("\n")
}
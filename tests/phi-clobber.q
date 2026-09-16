import io

main() {
	a := 3
	b := 3
	r := 7
	c := 7

	if a == b {
		c = b
	}

	r = (c >> 4) - (4 % 2) + c
	io.writeLine(r)
}
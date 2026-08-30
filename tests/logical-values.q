import io

main() {
	x := true && false
	y := true || false
	z := false && true
	w := false || true
	io.writeLine(boolToInt(x))
	io.writeLine(boolToInt(y))
	io.writeLine(boolToInt(z))
	io.writeLine(boolToInt(w))

	a := 1 < 2
	b := 2 < 1
	io.writeLine(boolToInt(a && b))
	io.writeLine(boolToInt(a || b))
	io.writeLine(boolToInt(a && !b))
	io.writeLine(boolToInt(a || !b))
}

boolToInt(x bool) -> int {
	if x {
		return 1
	}

	return 0
}
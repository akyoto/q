import io

main() {
	x := true && false
	y := true || false
	z := false && true
	w := false || true
	io.writeLine(x)
	io.writeLine(y)
	io.writeLine(z)
	io.writeLine(w)

	a := 1 > 0
	b := 1 < 0
	io.writeLine(a && b)
	io.writeLine(a || b)
	io.writeLine(a && !b)
	io.writeLine(a || !b)
}
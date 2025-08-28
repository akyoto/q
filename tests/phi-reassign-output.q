import io

main() {
	a := 0
	b := 1
	c := 2
	d := 3
	e := 4

	if true {
		a += 1
		b += 1
		c += 1
		d += 1
		e += 1
	}

	io.write(a)
	io.write(b)
	io.write(c)
	io.write(d)
	io.write(e)
}
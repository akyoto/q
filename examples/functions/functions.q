import sys

main() {
	a = add(1, 2)
	b = add(3, 4)
	c = add(a, b)
	show(c)

	d = sub(50, 10)
	e = sub(40, 10)
	f = sub(d, e)
	show(f)

	g = mul(1, 1)
	h = mul(2, 5)
	i = mul(g, h)
	show(i)

	j = div(1000, 10)
	k = div(100, 10)
	l = div(j, k)
	show(l)
}

add(a, b) {
	return a + b
}

sub(a, b) {
	return a - b
}

mul(a, b) {
	return a * b
}

div(a, b) {
	return a / b
}

show(num) {
	sys.write(1, "123456789\n", num)
}

import sys

# main is the entry point for our program.
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

# add adds two numbers.
add(a, b) {
	return a + b
}

# sub subtracts two numbers.
sub(a, b) {
	return a - b
}

# mul multiplies two numbers.
mul(a, b) {
	return a * b
}

# div divides two numbers.
div(a, b) {
	return a / b
}

# show shows a number on the console.
# Printing integers to the console isn't implemented yet,
# so we need to use some hacks to check the contents of integers.
show(num) {
	sys.write(1, "123456789\n", num)
}

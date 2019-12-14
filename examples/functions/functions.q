import sys

main() {
	let a = add(1, 2)
	let b = add(3, 4)
	let c = add(a, b)
	show(c)

	let d = sub(50, 10)
	let e = sub(40, 10)
	let f = sub(d, e)
	show(f)

	let g = mul(1, 1)
	let h = mul(2, 5)
	let i = mul(g, h)
	show(i)

	let j = div(1000, 10)
	let k = div(100, 10)
	let l = div(j, k)
	show(l)
}

# add adds two numbers.
add(a Int, b Int) -> Int {
	return a + b
}

# sub subtracts two numbers.
sub(a Int, b Int) -> Int {
	return a - b
}

# mul multiplies two numbers.
mul(a Int, b Int) -> Int {
	return a * b
}

# div divides two numbers.
div(a Int, b Int) -> Int {
	return a / b
}

# show shows a number on the console.
# Printing integers to the console isn't implemented yet,
# so we need to use some hacks to check the contents of integers.
show(num Int) {
	sys.write(1, "123456789\n", num)
}

Outer {
	a Inner
	b Inner
}

Inner {
	x int
	y int
}

main() {
	o := Outer{
		a: Inner{x: 1, y: 2},
		b: Inner{x: 3, y: 4},
	}

	f(o)
}

f(o Outer) {
	assert o.a.x == 1
	assert o.a.y == 2
	assert o.b.x == 3
	assert o.b.y == 4
}
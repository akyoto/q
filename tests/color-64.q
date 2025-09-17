main() {
	f(Color{r: 1, g: 2, b: 3, a: 4})
}

f(c Color) {
	assert c.r == 1
	assert c.g == 2
	assert c.b == 3
	assert c.a == 4
}

Color {
	r uint64
	g uint64
	b uint64
	a uint64
}
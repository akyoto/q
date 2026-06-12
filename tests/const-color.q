Color {
	r uint8
	g uint8
	b uint8
	a uint8
}

const {
	red = Color{r: 255, g: 0, b: 0, a: 255}
	green = Color{r: 0, g: 255, b: 0, a: 255}
	blue = Color{r: 0, g: 0, b: 255, a: 255}
}

main() {
	assertRed(red)
	assertGreen(green)
	assertBlue(blue)
}

assertRed(c Color) {
	assert c.r == 255 && c.a == 255
}

assertGreen(c Color) {
	assert c.g == 255 && c.a == 255
}

assertBlue(c Color) {
	assert c.b == 255 && c.a == 255
}
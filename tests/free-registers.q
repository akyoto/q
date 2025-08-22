Color {
	r byte
	g byte
	b byte
	a byte
}

main() {
	f(Color{r: 1, g: 2, b: 3, a: 4})
	f(Color{r: 5, g: 6, b: 7, a: 8})
	f(Color{r: 9, g: 10, b: 11, a: 12})
	f(Color{r: 13, g: 14, b: 15, a: 16})
	f(Color{r: 1, g: 2, b: 3, a: 4})
	f(Color{r: 5, g: 6, b: 7, a: 8})
	f(Color{r: 9, g: 10, b: 11, a: 12})
	f(Color{r: 13, g: 14, b: 15, a: 16})
}

f(_ Color) {}
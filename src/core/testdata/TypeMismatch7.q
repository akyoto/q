main() {
	store(42)
}

store(p *byte) {
	[p] = 'A'
}
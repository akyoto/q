main() {
	x := acquire()
	y := x
	free(x)
	free(y)
}

acquire() -> !int { return 1 }
free(_ !int) {}
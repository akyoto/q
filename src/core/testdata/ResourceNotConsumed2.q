main() {
	x := acquire()
	x = acquire()
	free(x)
}

acquire() -> !int { return 1 }
free(_ !int) {}
main() {
	x := acquire()
	use(x)
	free(x)
}

acquire() -> !int { return 1 }
use(_ int) {}
free(_ !int) {}
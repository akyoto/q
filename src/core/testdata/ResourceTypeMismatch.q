main() {
	x := acquire()
	use(x)
	free(x)
}

acquire() -> !int { return 1 }
use(a int) { free(a) }
free(_ !int) {}
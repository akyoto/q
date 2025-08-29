main() {
	x := acquire()

	if true {
		use(x)
	} else {
		free(x)
	}
}

acquire() -> !int { return 1 }
use(_ int) {}
free(_ !int) {}
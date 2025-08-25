main() {
	x := acquire()
	use(x)
}

acquire() -> !int { return 1 }
use(_ int) {}
main() {
	value, err := f()
	use(value)
}

f() -> (int, error) {
	return 0, 0
}

use(_ int) {}
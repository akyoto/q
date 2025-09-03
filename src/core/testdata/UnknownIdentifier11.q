main() {
	value, err := f()

	if err == 0 {
		return
	}

	use(value)
}

f() -> (int, error) { return 0, 0 }
use(_ int) {}
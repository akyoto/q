main() {
	value, err := f()

	if err != 0 {
		use(value)
	}
}

f() -> (int, error) { return 0, 0 }
use(_ int) {}
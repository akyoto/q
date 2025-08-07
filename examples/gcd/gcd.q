import io

main() {
	io.writeInt(gcd(1071, 462))
}

gcd(a int, b int) -> int {
	loop {
		switch {
			a == b { return a }
			a > b  { a -= b }
			_      { b -= a }
		}
	}
}
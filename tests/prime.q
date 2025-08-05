import io

main() {
	for i := 2..15 {
		if isPrime(i) {
			if i != 2 {
				io.write(" ")
			}

			for 0..i {
				io.write(".")
			}
		}
	}
}

isPrime(x int) -> bool {
	if x == 2 {
		return true
	}

	if x % 2 == 0 {
		return false
	}

	i := 3

	loop {
		if i * i > x {
			return true
		}

		if x % i == 0 {
			return false
		}

		i += 2
	}
}
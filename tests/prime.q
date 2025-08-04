import io

main() {
	for i := 2..15 {
		if isPrime(i) == 1 {
			if i != 2 {
				io.write(" ")
			}

			for 0..i {
				io.write(".")
			}
		}
	}
}

isPrime(x int) -> int {
	if x == 2 {
		return 1
	}

	if x % 2 == 0 {
		return 0
	}

	i := 3

	loop {
		if i * i > x {
			return 1
		}

		if x % i == 0 {
			return 0
		}

		i += 2
	}
}
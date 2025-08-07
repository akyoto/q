import io

main() {
	fizzbuzz(15)
}

fizzbuzz(n int) {
	x := 1

	loop {
		switch {
			x % 15 == 0 { io.write("FizzBuzz") }
			x % 5 == 0  { io.write("Buzz") }
			x % 3 == 0  { io.write("Fizz") }
			_           { io.writeInt(x) }
		}

		x += 1

		if x > n {
			return
		}

		io.write(" ")
	}
}
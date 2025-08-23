import io

main() {
	fizzbuzz(15)
}

fizzbuzz(n int) {
	loop x := 1 .. n+1 {
		switch {
			x % 15 == 0 { io.write("FizzBuzz") }
			x % 5 == 0  { io.write("Buzz") }
			x % 3 == 0  { io.write("Fizz") }
			_           { io.write(x) }
		}

		if x != n {
			io.write(" ")
		}
	}
}
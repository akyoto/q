import io

main() {
	n := 15

	loop i := 1..n+1 {
		fizzbuzz(i)

		if i != n {
			io.write(" ")
		}
	}
}

fizzbuzz(x int) {
	switch {
		x % 15 == 0 { io.write("FizzBuzz") }
		x % 5 == 0  { io.write("Buzz") }
		x % 3 == 0  { io.write("Fizz") }
		_           { io.write(x) }
	}
}
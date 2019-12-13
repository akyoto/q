import sys
import mem

struct Point {
	x Int
	y Int
}

main() {
	#let p = Point()
	printInt(123)
}

printInt(n Int) {
	let x = n

	if x > 9 {
		let a = x / 10
		#x = x - 10 * a
		printInt(a)
	}

	let buffer = mem.allocate(1)
	store(buffer, 0, 1, 48)
	#buffer[0] = 48
	sys.write(1, buffer, 1)
	mem.free(buffer, 1)
}

import sys

main() {
	# Repeat 3 times
	for 0..3 {
		print("Hello")
	}

	# Repeat 6 times and assign loop counter to 'i'
	for i := 0..6 {
		sys.write(1, "Hello", i)
		sys.write(1, "\n", 1)
	}
}

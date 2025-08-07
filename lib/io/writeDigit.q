writeDigit(n int) {
	switch {
		n == 0 { write("0") }
		n == 1 { write("1") }
		n == 2 { write("2") }
		n == 3 { write("3") }
		n == 4 { write("4") }
		n == 5 { write("5") }
		n == 6 { write("6") }
		n == 7 { write("7") }
		n == 8 { write("8") }
		n == 9 { write("9") }
	}
}
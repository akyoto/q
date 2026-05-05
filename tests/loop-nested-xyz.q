import io

main() {
	loop x := 0..3 {
		loop y := 0..3 {
			loop z := 0..3 {
				io.write(x)
				io.write(",")
				io.write(y)
				io.write(",")
				io.write(z)
				io.write("\n")
			}
		}
	}
}
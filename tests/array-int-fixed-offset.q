import io

main() {
	x := new(int, 8192)
	fill(x)
	show(x)
}

fill(x []int) {
	x[127] = 127
	x[128] = 128
	x[4095] = 4095
	x[4096] = 4096
}

show(x []int) {
	io.write(x[127])
	io.write("\n")
	io.write(x[128])
	io.write("\n")
	io.write(x[4095])
	io.write("\n")
	io.write(x[4096])
	io.write("\n")
}
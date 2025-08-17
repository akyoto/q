import io

main() {
	s := hello()
	io.write(s)
}

hello() -> string {
	return "Hello\n"
}
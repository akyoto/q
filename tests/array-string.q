import io

main() {
	s := new(string, 2)
	s[0] = "Hello"
	s[1] = "World"
	io.write(s[0])
	io.write(s[1])
}
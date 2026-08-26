import io

main() {
	io.write("f: ")
	io.write(f())
}

f() -> int {
	pre := 0
	post := 0

	loop {
		pre += 1

		if pre == 2 {
			loop.stop()
		}

		post += 1
	}

	return post
}
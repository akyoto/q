import io

main() {
	line := new(byte, 80)

	loop row := 0..40 {
		loop col := 0..80 {
			cx := -160000 + col * 1250
			cy := -50000 + row * 2500
			line[col] = glyph(mandel(cx, cy))
		}

		io.writeLine(line)
	}
}

mandel(cx int, cy int) -> int {
	x := cx
	y := cy

	loop i := 0..400 {
		nx := (x * x - y * y) / 100000 + cx
		ny := (x * y * 2) / 100000 + cy
		x = nx
		y = ny

		if x * x + y * y > 40000000000 {
			return i
		}
	}

	return -1
}

glyph(n int) -> byte {
	if n < 0 {
		return '#'
	}

	bucket := n / 20

	if bucket > 9 {
		bucket = 9
	}

	return " .:-=+*#%@"[bucket]
}
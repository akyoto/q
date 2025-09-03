import io
import mem

readFile(path string) -> (!string, error) {
	newPath := mem.alloc(path.len + 1)

	loop i := 0..path.len {
		newPath[i] = path[i]
	}

	fd, err := open(newPath.ptr, 0, 0)
	mem.free(newPath)

	if err != 0 {
		return "", err
	}

	fileSize, err := size(fd)

	if err != 0 {
		close(fd)
		return "", err
	}

	buffer := mem.alloc(fileSize)
	pos := 0

	loop {
		n := io.readFrom(fd, buffer[pos..])

		if n <= 0 {
			close(fd)
			return buffer, 0
		}

		pos += n
	}
}
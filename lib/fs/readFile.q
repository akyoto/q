import io
import mem

readFile(path string) -> !string {
	newPath := mem.alloc(path.len + 1)

	loop i := 0..path.len {
		newPath[i] = path[i]
	}

	fd := open(newPath.ptr, 0, 0)
	mem.free(newPath)
	fileSize := size(fd)

	if fileSize < 512 {
		fileSize = 512
	}

	buffer := mem.alloc(fileSize)
	pos := 0

	loop {
		n := io.readFrom(fd, buffer[pos..])

		if n <= 0 {
			close(fd)
			return buffer
		}

		pos += n
	}
}
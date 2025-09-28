import c
import io
import mem

readFile(path string) -> (!string, error) {
	cpath := c.string(path)
	fd, err := open(cpath.ptr, 0, 0)
	mem.free(cpath)

	if err != 0 {
		return "", err
	}

	fileSize, err := size(fd)

	if err != 0 {
		close(fd)
		return "", err
	}

	if fileSize == 0 {
		close(fd)
		return "", -1
	}

	buffer := mem.alloc(fileSize)
	pos := 0

	loop {
		n, err := io.readFrom(fd, buffer[pos..])

		if err != 0 {
			mem.free(buffer)
			close(fd)
			return "", err
		}

		pos += n

		if pos >= buffer.len {
			close(fd)
			return buffer, 0
		}
	}
}
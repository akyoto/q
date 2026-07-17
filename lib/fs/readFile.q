import c
import io

readFile(path string) -> (!string, error) {
	cpath := c.string(path)
	fd, err := openRead(cpath.ptr)
	delete(cpath)

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

	buffer := new(byte, fileSize)
	pos := 0

	loop {
		n, err := io.readFrom(fd, buffer[pos..])

		if err != 0 {
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
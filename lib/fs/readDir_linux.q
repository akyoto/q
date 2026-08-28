import c
import io

readDir(path string) -> error {
	cPath := c.string(path)
	fd, err := openRead(cPath.ptr)
	delete(cPath)

	if err != 0 {
		return err
	}

	buffer := new(byte, 512)

	loop {
		n := syscall(_getdents64, fd, buffer.ptr, buffer.len) as int

		if n <= 0 {
			close(fd)
			return n
		}

		ptr := buffer.ptr
		end := buffer.ptr + n

		loop {
			assert ptr < end
			cName := ptr + 19
			cLength := c.length(cName)
			assert cLength <= 255
			name := string{ptr: cName, len: cLength}
			io.writeLine(name)
			dirent := ptr as *Dirent64
			ptr += dirent.d_reclen

			if ptr == end {
				loop.stop()
			}
		}
	}
}
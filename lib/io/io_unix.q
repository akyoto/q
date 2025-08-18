write(buffer string) -> (written int) {
	return syscall(_write, 1, buffer.ptr, buffer.len)
}

read(buffer string) -> (read int) {
	return syscall(_read, 0, buffer.ptr, buffer.len)
}
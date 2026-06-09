futex(address *uint32, op int, value int, timeout *any|nil, address2 *uint32|nil, value3 uint32) -> int {
	return syscall(_futex, address, op, value, timeout, address2, value3)
}
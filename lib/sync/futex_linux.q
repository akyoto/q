import time

futex_wake(address *any, mask uint, count int32, flags uint32) -> error {
	return syscall(_futex_wake, address, mask, count, flags)
}

futex_wait(address *any, value uint, mask uint, flags uint32, timeout *time.Timespec|nil, clockid int32) -> error {
    return syscall(_futex_wait, address, value, mask, flags, timeout, clockid)
}

futex_waitv(waiters *FutexWaitv, count uint32, flags uint32, timeout *time.Timespec|nil, clockid int32) -> error {
	return syscall(_futex_waitv, waiters, count, flags, timeout, clockid)
}
import time

wait(address *uint32, value uint32) -> error {
	return futex_wait(address, value as uint64, 0xFFFFFFFF, FUTEX2_SIZE_U32, 0, time.monotonic)
}
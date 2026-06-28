import time

wake(state *uint32, count int32) -> error {
	return futex_wake(state, 0xFFFFFFFF, count, FUTEX2_SIZE_U32 | FUTEX2_PRIVATE)
}

wait(address *uint32, value uint32) -> error {
	return futex_wait(address, value as uint64, 0xFFFFFFFF, FUTEX2_SIZE_U32 | FUTEX2_PRIVATE, 0, time.monotonic)
}
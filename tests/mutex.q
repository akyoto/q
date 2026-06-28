import sync

main() {
	mutex := new(uint32)

	loop 0..10 {
		// TODO: Spawn multiple threads once threading works on all platforms
		work(mutex)
	}
}

work(mutex *uint32) {
	sync.lock(mutex)
	loop 0..100000 {}
	sync.unlock(mutex)
}
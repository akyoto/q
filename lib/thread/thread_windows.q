create(func any) -> int {
	return kernel32.CreateThread(0, STACK_SIZE, func, 0)
}

extern {
	kernel32 {
		CreateThread(attributes int, stackSize int, address any, parameter int) -> int
	}
}
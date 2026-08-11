main() {
	stack := _cpu.sp as *int
	assert stack != 0
	assert [stack] != 0
}
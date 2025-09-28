main() {
	stack := asm.sp as *int
	assert stack != 0
	assert [stack] != 0
}
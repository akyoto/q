main() {
	sysExit = 60
	exitCode = (sysExit + 70 - 80) * 2
	syscall(sysExit, exitCode)
}

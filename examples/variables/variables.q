main() {
	sysExit = 20 + 40
	exitCode = sysExit + 15 + sysExit + 65
	syscall(sysExit, exitCode)
}

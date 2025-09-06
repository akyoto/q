ProcessInformation {
	process *any
	thread *any
	processId uint32
	threadId uint32
}

StartupInfo {
	size uint32
	reserved *uint16
	desktop *uint16
	title *uint16
	x uint32
	y uint32
	width uint32
	height uint32
	xCountChars uint32
	yCountChars uint32
	fillAttribute uint32
	flags uint32
	showWindow uint16
	reserved2Size uint16
	reserved2 *byte
	stdIn *any
	stdOut *any
	stdError *any
}
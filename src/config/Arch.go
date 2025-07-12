package config

// Arch is a CPU architecture.
type Arch uint8

const (
	UnknownArch Arch = iota
	ARM
	X86
)
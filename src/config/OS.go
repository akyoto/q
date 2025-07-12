package config

// OS is an operating system.
type OS uint8

const (
	UnknownOS OS = iota
	Linux
	Mac
	Windows
)
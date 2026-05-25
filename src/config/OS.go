package config

// OS is an operating system.
type OS uint8

const (
	UnknownOS OS = iota
	Linux
	Mac
	Windows
)

// String returns the lowercase name of the operating system.
func (os OS) String() string {
	switch os {
	case Linux:
		return "linux"
	case Mac:
		return "mac"
	case Windows:
		return "windows"
	default:
		return "unknown"
	}
}